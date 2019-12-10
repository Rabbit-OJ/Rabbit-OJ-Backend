package initialize

import (
	"Rabbit-OJ-Backend/services/judger"
	"Rabbit-OJ-Backend/utils"
	"strconv"
)

func CheckTestCase()  {
	storage := judger.ReadStorageFile()
	newStorage := make([]judger.Storage, 0)

	for _, item := range storage {
		valid := true

		for i := uint32(1); i <= item.DatasetCount; i++ {
			inputFilePath, err := utils.JudgeFilePath(item.Tid, item.Version, strconv.FormatUint(uint64(i), 10), "in")
			if err != nil {
				valid = false
			}
			if !utils.Exists(inputFilePath) {
				valid = false
			}

			outputFilePath, err := utils.JudgeFilePath(item.Tid, item.Version, strconv.FormatUint(uint64(i), 10), "out")
			if err != nil {
				valid = false
			}
			if !utils.Exists(outputFilePath) {
				valid = false
			}
		}

		if valid {
			newStorage = append(newStorage, item)
		}
	}

	if err := judger.SaveStorageFile(newStorage); err != nil {
		panic(err)
	}
}

func PruneTestCase() {
	// todo: delete not referenced object
}