package question

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func Case(tid string, response *protobuf.TestCaseResponse) error {
	judgeObj, err := JudgeInfo(tid)
	if err != nil {
		return err
	}

	count := judgeObj.DatasetCount
	*response = protobuf.TestCaseResponse{
		Tid:  tid,
		Case: make([]*protobuf.Case, 0, count),
	}

	for i := uint32(0); i < count; i++ {
		currentCase := &protobuf.Case{}

		inFile, errIn := readCaseFile(tid, i, "in")
		if errIn != nil {
			fmt.Println(errIn)
			continue
		}

		outFile, errOut := readCaseFile(tid, i, "out")
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

func readCaseFile(tid string, caseId uint32, mode string) ([]byte, error) {
	file, err := os.Open(utils.CaseFilePath(tid, caseId, mode))
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
