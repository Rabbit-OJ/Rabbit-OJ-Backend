package question

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func Case(tid string, response *protobuf.TestCaseResponse) error {
	judgeObj, err := JudgeInfo(tid)
	if err != nil {
		return err
	}

	count := judgeObj.DatasetCount
	*response = protobuf.TestCaseResponse{
		Tid:     tid,
		Case:    make([]*protobuf.Case, 0, count),
		Version: strconv.FormatUint(uint64(judgeObj.Version), 10),
	}

	for i := uint32(1); i <= count; i++ {
		currentCase := &protobuf.Case{}

		inFile, errIn := readCaseFile(
			tid,
			strconv.FormatUint(uint64(judgeObj.Version), 10),
			strconv.FormatUint(uint64(i), 10), "in")
		if errIn != nil {
			fmt.Println(errIn)
			continue
		}

		outFile, errOut := readCaseFile(
			tid,
			strconv.FormatUint(uint64(judgeObj.Version), 10),
			strconv.FormatUint(uint64(i), 10), "out")
		if errOut != nil {
			fmt.Println(errOut)
			continue
		}

		currentCase.Stdin, currentCase.Stdout = inFile, outFile
		response.Case = append(response.Case, currentCase)
	}

	if len(response.Case) == 0 {
		return errors.New("no valid cases found")
	}

	return nil
}

func readCaseFile(tid, version, caseId, caseType string) ([]byte, error) {
	filePath, err := utils.JudgeFilePath(tid, version, caseId, caseType)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer func() { _ = file.Close() }()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
