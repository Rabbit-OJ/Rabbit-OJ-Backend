package judger

import (
	"Rabbit-OJ-Backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	STATUS_OK  = "OK"
	STATUS_TLE = "TLE"
	STATUS_MLE = "MLE"
	STATUS_RE  = "RE"
)

type TestResult struct {
	CaseId    int64
	Status    string
	TimeUsed  uint32
	SpaceUsed uint32
}

func Tester() {
	// <-- step1 : validate
	testCaseCount, err := strconv.ParseInt(os.Getenv("CASE_COUNT"), 10, 32)

	if err != nil {
		panic(err)
	}

	if testCaseCount <= 0 {
		panic(errors.New("invalid test case"))
	}

	for i := int64(1); i <= testCaseCount; i++ {
		if !utils.Exists(utils.DockerCasePath(i)) {
			panic(errors.New(fmt.Sprintf("Case #%d doesn't exist", i)))
		}
	}

	file, err := os.Create("/result/info.json")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = file.Close()
	}()

	// <-- step2 : get_result
	testResult := make([]TestResult, testCaseCount)
	for i := int64(1); i <= testCaseCount; i++ { // TODO: LIMIT MEMORY & judge

	}

	// <-- step3 : write info
	result, err := json.Marshal(testResult)
	if err != nil {
		panic(err)
	}

	if _, err := file.Write(result); err != nil {
		panic(err)
	}
}
