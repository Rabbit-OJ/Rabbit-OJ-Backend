package storage

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/storage/rpc"
	"Rabbit-OJ-Backend/utils/files"
	"encoding/json"
	"fmt"
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
