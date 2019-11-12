package judger

import (
	"Rabbit-OJ-Backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	StatusOK  = "OK"
	StatusTLE = "TLE"
	StatusMLE = "MLE"
	StatusRE  = "RE"
)

type TestResult struct {
	CaseId    int64
	Status    string
	TimeUsed  uint32
	SpaceUsed uint32
}

func max(a, b int64) int64 {
	if a < b {
		return b
	} else {
		return a
	}
}

func TestOne(testResult []TestResult, i, timeLimit, spaceLimit int64) {
	cmd := exec.Command("/compile/code.o")
	peakMemory := int64(0)

	if err := cmd.Start(); err != nil {
		testResult[i-1].Status = StatusRE
	}

	in, err := os.OpenFile("E:/test.in", os.O_RDONLY, 0644)
	if err != nil {
		testResult[i-1].Status = StatusRE
		return
	}
	defer func() {
		_ = in.Close()
	}()

	out, err := os.OpenFile("E:/test.out", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = out.Close()
	}()

	cmd.Stdin = in
	cmd.Stdout = out

	if err := cmd.Start(); err != nil {
		testResult[i-1].Status = StatusRE
		return
	}
	startTime := time.Now()

	errChan, successChan, memoryMonitorChan, memoryMonitorCloseChan := make(chan error), make(chan bool), make(chan bool), make(chan bool)
	go func() {
		err := cmd.Wait()
		if err != nil {
			errChan <- err
		} else {
			successChan <- true
		}
	}()
	go func(pid int) {
		for {
			select {
			case <-memoryMonitorCloseChan:
				return
			default:
				stat, err := utils.GetStat(pid)
				if err == nil {
					peakMemory = max(peakMemory,
						int64(stat.Memory/1024/1024),
					)

					if peakMemory >= spaceLimit {
						memoryMonitorChan <- true
					}
				}
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(cmd.Process.Pid)

	select {
	case <-memoryMonitorChan:
		testResult[i-1].Status = StatusMLE
		testResult[i-1].TimeUsed = uint32(timeLimit)
		_ = cmd.Process.Kill()
	case <-successChan:
		testResult[i-1].Status = StatusOK
		usedTime := time.Since(startTime)
		testResult[i-1].TimeUsed = uint32(usedTime.Milliseconds())
		testResult[i-1].SpaceUsed = uint32(peakMemory)
	case <-time.After(time.Duration(timeLimit) * time.Millisecond):
		testResult[i-1].Status = StatusTLE
		testResult[i-1].TimeUsed = uint32(timeLimit)
		testResult[i-1].SpaceUsed = uint32(peakMemory)
		_ = cmd.Process.Kill()
	case <-errChan:
		testResult[i-1].Status = StatusRE
	}

	memoryMonitorCloseChan <- true
}

func Tester() {
	// <-- step1 : validate
	testCaseCount, err := strconv.ParseInt(os.Getenv("CASE_COUNT"), 10, 32)
	if err != nil {
		panic(err)
	}
	timeLimit, err := strconv.ParseInt(os.Getenv("TIME_LIMIT"), 10, 32)
	if err != nil {
		panic(err)
	}
	spaceLimit, err := strconv.ParseInt(os.Getenv("SPACE_LIMIT"), 10, 32)
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
	for i := int64(1); i <= testCaseCount; i++ {
		TestOne(testResult, i, timeLimit, spaceLimit)
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
