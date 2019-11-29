package judger

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

type CollectedStdout struct {
	Stdout      string
	RightStdout string
}

func Scheduler(request *protobuf.JudgeRequest) error {
	sid := request.Sid

	fmt.Println("========START JUDGE ========")
	fmt.Println("[Scheduler] Received judge request " + sid)
	// init path
	currentPath, err := utils.SubmissionGenerateDirWithMkdir(sid)
	if err != nil {
		return err
	}

	outputPath, err := utils.JudgeGenerateOutputDirWithMkdir(currentPath)
	if err != nil {
		return err
	}

	codePath := fmt.Sprintf("%s/%s.code", currentPath, sid)

	jsonPath := codePath + ".json"
	casePath, err := utils.JudgeCaseDir(request.Tid, request.Version)
	if err != nil {
		return err
	}

	compileInfo, ok := utils.CompileObject[request.Language]
	if !ok {
		return errors.New("language doesn't support")
	}

	fmt.Println("[Scheduler] Init test cases")
	// get case
	storage, err := InitTestCase(request.Tid, request.Version)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// compile
	fmt.Println("[Scheduler] Start Compile")
	if err := Compiler(codePath, request.Code, &compileInfo); err != nil {
		fmt.Println("[Scheduler] CE", err)
		return callbackAllError("CE", sid, storage)
	}
	fmt.Println("[Scheduler] Compile OK")

	// run
	fmt.Println("[Scheduler] Start Runner")
	if err := Runner(
		codePath,
		&compileInfo,
		strconv.FormatUint(uint64(storage.DatasetCount), 10),
		strconv.FormatUint(uint64(request.TimeLimit), 10),
		strconv.FormatUint(uint64(request.SpaceLimit), 10),
		casePath,
		outputPath); err != nil {

		fmt.Println("RE", err)
		return callbackAllError("RE", sid, storage)
	}
	fmt.Println("[Scheduler] Runner OK")

	fmt.Println("[Scheduler] Reading result")
	jsonFileByte, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return callbackAllError("RE", sid, storage)
	}

	var testResultArr []models.TestResult
	if err := json.Unmarshal(jsonFileByte, &testResultArr); err != nil {
		return callbackAllError("RE", sid, storage)
	}

	// collect std::out
	fmt.Println("[Schedule] Collecting stdout " + sid)
	allStdin := make([]CollectedStdout, storage.DatasetCount)
	for i := uint32(1); i <= storage.DatasetCount; i++ {

		path, err := utils.JudgeFilePath(
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
	fmt.Println("Judging stdout " + sid)
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
	fmt.Println("[Scheduler] Calling back results")
	go callbackWebSocket(sid)
	if err := callbackSuccess(
		sid,
		resultList); err != nil {
		fmt.Println(err)
		return err
	}
	// todo: clear cache

	fmt.Println("[Scheduler] Finish " + sid)
	return nil
}
