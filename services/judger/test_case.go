package judger

import (
	"Rabbit-OJ-Backend/protobuf"
	"Rabbit-OJ-Backend/services/rpc"
	"Rabbit-OJ-Backend/utils"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Storage struct {
	Tid          string
	Version      string
	DatasetCount uint32
}

type StorageFile = []Storage

func ReadStorageFile() []Storage {
	storageFilePath, err := utils.StorageFilePath()
	if err != nil {
		panic(err)
	}

	storage := make([]Storage, 0)

	if !utils.Exists(storageFilePath) {
		return storage
	}

	rawJson, err := ioutil.ReadFile(storageFilePath)
	if err != nil {
		return storage
	}

	if err := json.Unmarshal(rawJson, storage); err != nil {
		return storage
	}

	return storage
}

func SaveStorageFile(storage []Storage) error {
	storageFilePath, err := utils.StorageFilePath()
	if err != nil {
		panic(err)
	}

	file, err := json.Marshal(storage)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(storageFilePath, file, 0644); err != nil {
		return err
	}

	return nil
}

func FetchTestCase(tid string) (*Storage, error) {
	request, response := &protobuf.TestCaseRequest{Tid: tid}, &protobuf.TestCaseResponse{}

	if err := rpc.DialCall("CaseService", "Case", request, response); err != nil {
		return nil, err
	}

	_, err := utils.JudgeDirPathWithMkdir(tid, response.Version)
	if err != nil {
		return nil, err
	}

	for index, item := range response.Case {
		inPath, err := utils.JudgeFilePath(tid, response.Version, strconv.Itoa(index), "in")
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(inPath, item.Stdin, 0644); err != nil {
			return nil, err
		}

		outPath, err := utils.JudgeFilePath(tid, response.Version, strconv.Itoa(index), "out")
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(outPath, item.Stdout, 0644); err != nil {
			return nil, err
		}
	}

	storage := ReadStorageFile()
	newStorage := &Storage{
		Tid:          tid,
		Version:      response.Version,
		DatasetCount: uint32(len(response.Case)),
	}
	storage = append(storage, *newStorage)
	return newStorage, SaveStorageFile(storage)
}

func InitTestCase(tid, version string) (*Storage, error) {
	storageFileContent := ReadStorageFile()

	containsFlag, storage := false, &Storage{}
	for _, item := range storageFileContent {
		if item.Tid == tid {
			storage = &item
			containsFlag = true
			break
		}
	}

	if (!containsFlag) || (containsFlag && storage.Version != version) {
		return FetchTestCase(tid)
	} else {
		return storage, nil
	}
}
