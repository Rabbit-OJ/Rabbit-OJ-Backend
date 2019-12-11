package judger

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/utils/files"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type CollectedStdout struct {
	Stdout      string
	RightStdout string
}

func Scheduler(request *protobuf.JudgeRequest) error {
	sid := request.Sid

	fmt.Printf("========START JUDGE(%s)======== \n", sid)
	fmt.Printf("(%s) [Scheduler] Received judge request \n", sid)

	startSchedule := time.Now()
	defer func() {
		fmt.Printf("(%s) [Scheduler] total cost : %d ms \n", sid, time.Since(startSchedule).Milliseconds())
	}()

	// initialize files
	currentPath, err := files.SubmissionGenerateDirWithMkdir(sid)
	if err != nil {
		return err
	}

	defer func() {
		fmt.Printf("(%s) [Scheduler] Cleaning files \n", sid)
		if config.Global.AutoRemove.Files {
			_ = os.RemoveAll(currentPath)
		}
	}()

	outputPath, err := files.JudgeGenerateOutputDirWithMkdir(currentPath)
	if err != nil {
		return err
	}

	codePath := fmt.Sprintf("%s/%s.code", currentPath, sid)
	casePath, err := files.JudgeCaseDir(request.Tid, request.Version)
	if err != nil {
		return err
	}

	compileInfo, ok := config.CompileObject[request.Language]
	if !ok {
		return errors.New("language doesn't support")
	}

	fmt.Printf("(%s) [Scheduler] Init test cases \n", sid)
	// get case
	storage, err := InitTestCase(request.Tid, request.Version)
	if err != nil {
		fmt.Printf("(%s) [Scheduler] Case Error %+v \n", sid, err)
		return err
	}

	if !compileInfo.NoBuild {
		// compile
		fmt.Printf("(%s) [Scheduler] Start Compile \n", sid)
		if err := Compiler(sid, codePath, request.Code, &compileInfo); err != nil {
			fmt.Printf("(%s) [Scheduler] CE %+v \n", sid, err)
			callbackAllError("CE", sid, storage)
		}
		fmt.Printf("(%s) [Scheduler] Compile OK \n", sid)

		fileStat, err := os.Stat(codePath + ".o")
		if err != nil || fileStat.Size() == 0 {
			fmt.Printf("(%s) [Scheduler] CE %+v \n", sid, err)
			callbackAllError("CE", sid, storage)
		}
	}

	// run
	fmt.Printf("(%s) [Scheduler] Start Runner \n", sid)
	if err := Runner(
		sid,
		codePath,
		&compileInfo,
		strconv.FormatUint(uint64(storage.DatasetCount), 10),
		strconv.FormatUint(uint64(request.TimeLimit), 10),
		strconv.FormatUint(uint64(request.SpaceLimit), 10),
		casePath,
		outputPath,
		request.Code); err != nil {

		fmt.Printf("(%s) [Scheduler] RE %+v \n", sid, err)
		callbackAllError("RE", sid, storage)
	}
	fmt.Printf("(%s) [Scheduler] Runner OK \n", sid)

	fmt.Printf("(%s) [Scheduler] Reading result \n", sid)
	jsonFileByte, err := ioutil.ReadFile(codePath + ".result")
	if err != nil {
		callbackAllError("RE", sid, storage)
	}

	var testResultArr []models.TestResult
	if err := json.Unmarshal(jsonFileByte, &testResultArr); err != nil {
		callbackAllError("RE", sid, storage)
	}

	// collect std::out
	fmt.Printf("(%s) [Schedule] Collecting stdout \n", sid)
	allStdin := make([]CollectedStdout, storage.DatasetCount)
	for i := uint32(1); i <= storage.DatasetCount; i++ {

		path, err := files.JudgeFilePath(
			storage.Tid,
			storage.Version,
			strconv.FormatUint(uint64(i), 10),
			"out")

		if err != nil {
			return err
		}

		stdoutByte, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		allStdin[i-1].RightStdout = string(stdoutByte)
	}

	for i := uint32(1); i <= storage.DatasetCount; i++ {
		path := fmt.Sprintf("%s/%d.out", outputPath, i)

		stdoutByte, err := ioutil.ReadFile(path)
		if err != nil {
			allStdin[i-1].Stdout = ""
		} else {
			allStdin[i-1].Stdout = string(stdoutByte)
		}
	}
	// judge std::out
	fmt.Printf("(%s) [Scheduler] Judging stdout \n", sid)
	resultList := make([]*protobuf.JudgeCaseResult, storage.DatasetCount)

	for index, item := range allStdin {
		testResult := &testResultArr[index]
		resultList[index] = &protobuf.JudgeCaseResult{}

		judgeResult := JudgeOneCase(testResult, item.Stdout, item.RightStdout, request.CompMode)
		resultList[index].Status = judgeResult.Status
		resultList[index].SpaceUsed = judgeResult.SpaceUsed
		resultList[index].TimeUsed = judgeResult.TimeUsed
	}
	// mq return result
	fmt.Printf("(%s) [Scheduler] Calling back results \n", sid)
	callbackSuccess(sid, resultList)

	fmt.Printf("(%s) [Scheduler] Finish \n", sid)
	return nil
}
