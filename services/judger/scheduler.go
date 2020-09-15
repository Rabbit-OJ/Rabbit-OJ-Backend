package judger

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/config"
	StorageService "Rabbit-OJ-Backend/services/storage"
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

func Scheduler(request *protobuf.JudgeRequest) (bool, error) {
	sid := request.Sid

	fmt.Printf("========START JUDGE(%d)======== \n", sid)
	fmt.Printf("(%d) [Scheduler] Received judge request \n", sid)

	startSchedule := time.Now()
	defer func() {
		fmt.Printf("(%d) [Scheduler] total cost : %d ms \n", sid, time.Since(startSchedule).Milliseconds())
	}()

	// initialize files
	currentPath, err := files.SubmissionGenerateDirWithMkdir(sid)
	if err != nil {
		return false, err
	}

	defer func() {
		fmt.Printf("(%d) [Scheduler] Cleaning files \n", sid)
		if config.Global.AutoRemove.Files {
			_ = os.RemoveAll(currentPath)
		}
	}()

	outputPath, err := files.JudgeGenerateOutputDirWithMkdir(currentPath)
	if err != nil {
		return false, err
	}

	codePath := fmt.Sprintf("%s/", currentPath)
	casePath, err := files.JudgeCaseDir(request.Tid, request.Version)
	if err != nil {
		return false, err
	}

	compileInfo, ok := config.CompileObject[request.Language]
	if !ok {
		return false, errors.New("language doesn't support")
	}

	fmt.Printf("(%d) [Scheduler] Init test cases \n", sid)
	// get case
	storage, err := StorageService.InitTestCase(request.Tid, request.Version)
	if err != nil {
		fmt.Printf("(%d) [Scheduler] Case Error %+v \n", sid, err)
		return false, err
	}

	if !compileInfo.NoBuild {
		// compile
		fmt.Printf("(%d) [Scheduler] Start Compile \n", sid)
		if err := Compiler(sid, codePath, request.Code, &compileInfo); err != nil {
			fmt.Printf("(%d) [Scheduler] CE %+v \n", sid, err)
			CallbackAllError("CE", sid, request.IsContest, storage)
			return true, err
		}
		fmt.Printf("(%d) [Scheduler] Compile OK \n", sid)
	}

	// run
	fmt.Printf("(%d) [Scheduler] Start Runner \n", sid)
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

		fmt.Printf("(%d) [Scheduler] RE %+v \n", sid, err)
		CallbackAllError("RE", sid, request.IsContest, storage)
		return true, err
	}
	fmt.Printf("(%d) [Scheduler] Runner OK \n", sid)

	fmt.Printf("(%d) [Scheduler] Reading result \n", sid)
	jsonFileByte, err := ioutil.ReadFile(codePath + "result.json")
	if err != nil {
		CallbackAllError("RE", sid, request.IsContest, storage)
		return true, err
	}

	var testResultArr []models.TestResult
	if err := json.Unmarshal(jsonFileByte, &testResultArr); err != nil || testResultArr == nil {
		CallbackAllError("RE", sid, request.IsContest, storage)
		return true, err
	}

	// collect std::out
	fmt.Printf("(%d) [Scheduler] Collecting stdout \n", sid)
	allStdin := make([]CollectedStdout, storage.DatasetCount)
	for i := uint32(1); i <= storage.DatasetCount; i++ {

		path, err := files.JudgeFilePath(
			storage.Tid,
			storage.Version,
			strconv.FormatUint(uint64(i), 10),
			"out")

		if err != nil {
			return true, err
		}

		stdoutByte, err := ioutil.ReadFile(path)
		if err != nil {
			return true, err
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
	fmt.Printf("(%d) [Scheduler] Judging stdout \n", sid)
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
	fmt.Printf("(%d) [Scheduler] Calling back results \n", sid)
	CallbackSuccess(sid, request.IsContest, resultList)

	fmt.Printf("(%d) [Scheduler] Finish \n", sid)
	return true, nil
}
