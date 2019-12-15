package judger

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/utils"
	"Rabbit-OJ-Backend/utils/files"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
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

func TestOne(
	testResult *models.TestResult,
	i, timeLimit, spaceLimit int64,
	execCommand string,
	execArgs []string) {

	cmd := exec.Command(execCommand, execArgs...)
	peakMemory := float64(0)

	in, err := os.OpenFile(files.DockerCasePath(i), os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
		testResult.Status = StatusRE
		return
	}
	defer func() {
		_ = in.Close()
	}()

	out, err := os.OpenFile(files.DockerOutputPath(i), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		_ = out.Close()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	errChan, memoryMonitorChan := make(chan error), make(chan bool)
	defer func() {
		close(errChan)
		close(memoryMonitorChan)

		cancel()
	}()

	cmd.Stdin, cmd.Stdout = in, out

	if err := cmd.Start(); err != nil {
		log.Println(err)
		testResult.Status = StatusRE
		return
	}
	startTime := time.Now()

	go func() {
		waitChan := make(chan error)

		go func() {
			err := cmd.Wait()
			waitChan <- err
			close(waitChan)
		}()

		select {
		case <-ctx.Done():
			return
		case ans := <-waitChan:
			errChan <- ans
		}
	}()

	go func(pid int) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				stat, err := utils.GetStat(pid)
				if err == nil {
					peakMemory = math.Max(peakMemory,
						stat.Memory/1024.0,
					)

					if peakMemory >= float64(spaceLimit) * 1024.0 {
						memoryMonitorChan <- true
					}
				}
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(cmd.Process.Pid)

	select {
	case <-memoryMonitorChan:
		testResult.Status = StatusMLE
		testResult.TimeUsed = uint32(timeLimit)
		_ = cmd.Process.Kill()
	case <-time.After(time.Duration(timeLimit) * time.Millisecond):
		testResult.Status = StatusTLE
		testResult.TimeUsed = uint32(timeLimit)
		testResult.SpaceUsed = peakMemory
		_ = cmd.Process.Kill()
	case err := <-errChan:
		usedTime := time.Since(startTime)

		if err != nil {
			log.Println(err)
			testResult.Status = StatusRE
		} else {
			testResult.Status = StatusOK
		}

		testResult.TimeUsed = uint32(usedTime.Milliseconds())
		testResult.SpaceUsed = peakMemory
	}
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

	// todo: optimistic ? can we believe the scheduler and do less routine ???
	if testCaseCount <= 0 {
		panic(errors.New("invalid test case"))
	}

	for i := int64(1); i <= testCaseCount; i++ {
		if !files.Exists(files.DockerCasePath(i)) {
			panic(errors.New(fmt.Sprintf("Case #%d doesn't exist", i)))
		}
	}

	execCommandRaw := os.Getenv("EXEC_COMMAND")
	if execCommandRaw == "" {
		panic(err)
	}
	var execCommandArr []string
	if err := json.Unmarshal([]byte(execCommandRaw), &execCommandArr); err != nil {
		panic(err)
	}
	execCommand, execArgs := execCommandArr[0], execCommandArr[1:]

	if len(execCommandArr) == 1 {
		if err := os.Chmod(execCommandArr[0], 0755); err != nil {
			log.Println(err)
		}
	}

	file, err := os.Create(files.DockerResultFile())
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	// <-- step2 : get_result
	testResult := make([]models.TestResult, testCaseCount)
	for i := int64(1); i <= testCaseCount; i++ {
		fmt.Printf("Test #%d Case...\n", i)
		TestOne(&testResult[i-1], i, timeLimit, spaceLimit, execCommand, execArgs)
	}

	// <-- step3 : write info
	result, err := json.Marshal(testResult)
	if err != nil {
		panic(err)
	}
	if _, err := file.Write(result); err != nil {
		panic(err)
	}

	os.Exit(0)
}
