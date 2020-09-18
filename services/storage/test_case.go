package storage

import (
	"Rabbit-OJ-Backend/models/protobuf"
	"Rabbit-OJ-Backend/services/storage/rpc"
	"Rabbit-OJ-Backend/utils/files"
	"encoding/json"
	"fmt"
	JudgerModels "github.com/Rabbit-OJ/Rabbit-OJ-Judger/models"
	"github.com/Rabbit-OJ/Rabbit-OJ-Judger/utils"
	"io/ioutil"
	"strconv"
	"sync"
)

var (
	fetchTestCaseMutex sync.Mutex
)

type Storage struct {
	Tid          uint32
	Version      string
	DatasetCount uint32
}

type StorageFile = []Storage

func ReadStorageFile() []Storage {
	storageFilePath, err := files.StorageFilePath()
	storage := make([]Storage, 0)

	if err != nil {
		fmt.Println(err)
		return storage
	}

	if !files.Exists(storageFilePath) {
		return storage
	}

	rawJson, err := ioutil.ReadFile(storageFilePath)
	if err != nil {
		fmt.Println(err)
		return storage
	}

	if err := json.Unmarshal(rawJson, &storage); err != nil {
		fmt.Println(err)
		return storage
	}

	return storage
}

func SaveStorageFile(storage []Storage) error {
	storageFilePath, err := files.StorageFilePath()
	if err != nil {
		fmt.Println(err)
	}

	file, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := ioutil.WriteFile(storageFilePath, file, 0644); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func FetchTestCase(tid uint32, storage []Storage) (*Storage, error) {
	fetchTestCaseMutex.Lock()
	defer fetchTestCaseMutex.Unlock()

	request, response := &protobuf.TestCaseRequest{Tid: tid}, &protobuf.TestCaseResponse{}

	fmt.Printf("[Test Case] Preparing rpc to fetch case %d \n", tid)
	if err := rpc.DialCall("CaseService", "Case", request, response); err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("[Test Case] Storage to localhost")
	if _, err := files.JudgeDirPathWithMkdir(tid, response.Version); err != nil {
		return nil, err
	}

	for index, item := range response.Case {
		inPath, err := files.JudgeFilePath(tid, response.Version, strconv.Itoa(index+1), "in")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		if err := ioutil.WriteFile(inPath, item.Stdin, 0644); err != nil {
			fmt.Println(err)
			return nil, err
		}

		outPath, err := files.JudgeFilePath(tid, response.Version, strconv.Itoa(index+1), "out")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		if err := ioutil.WriteFile(outPath, item.Stdout, 0644); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	fmt.Println("[Test Case] Writing to index")
	newStorage := &Storage{
		Tid:          tid,
		Version:      response.Version,
		DatasetCount: uint32(len(response.Case)),
	}

	fmt.Println("[Test Case] Fetch OK")
	return newStorage, SaveStorageFile(append(storage, *newStorage))
}

func InitTestCase(tid uint32, version string) (*Storage, error) {
	fmt.Println("[Test Case] Reading index")
	storageFileContent := ReadStorageFile()

	fmt.Println("[Test Case] checking valid status")
	containsFlag, storage := false, &Storage{}
	for _, item := range storageFileContent {
		if item.Tid == tid {
			storage = &item
			containsFlag = true
			break
		}
	}

	if (!containsFlag) || (containsFlag && storage.Version != version) {
		fmt.Println("[Test Case] Check failed, Start Fetch")
		return FetchTestCase(tid, storageFileContent)
	} else {
		fmt.Println("[Test Case] Check OK")
		return storage, nil
	}
}

func GetTestCase(tid uint32, version string) ([]*JudgerModels.TestCaseType, error) {
	storage, err := InitTestCase(tid, version)
	if err != nil {
		return nil, err
	}

	datasetCount := int(storage.DatasetCount)
	resp := make([]*JudgerModels.TestCaseType, datasetCount)
	for i := 1; i <= datasetCount; i++ {
		stdoutPath, err := utils.JudgeFilePath(
			tid,
			version,
			strconv.FormatUint(uint64(i), 10),
			"out")
		if err != nil {
			return nil, err
		}

		stdoutByte, err := ioutil.ReadFile(stdoutPath)
		if err != nil {
			return nil, err
		}

		stdinPath, err := utils.JudgeFilePath(
			tid,
			version,
			strconv.FormatUint(uint64(i), 10),
			"in")
		if err != nil {
			return nil, err
		}

		stdinByte, err := ioutil.ReadFile(stdinPath)
		if err != nil {
			return nil, err
		}

		resp[i-1] = &JudgerModels.TestCaseType{
			Id:         i,
			StdinPath:  stdinPath,
			StdoutPath: stdoutPath,
			Stdin:      stdinByte,
			Stdout:     stdoutByte,
		}
	}
	return resp, nil
}
