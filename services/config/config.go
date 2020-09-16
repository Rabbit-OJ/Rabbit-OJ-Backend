package config

import (
	"Rabbit-OJ-Backend/utils/files"
	"encoding/json"
	JudgerModels "github.com/Rabbit-OJ/Rabbit-OJ-Judger/models"
	"io/ioutil"
)

var (
	Global *GlobalConfigType
)

type GlobalConfigType struct {
	MySQL  string                        `json:"mysql"`
	Debug  debug                         `json:"debug"`
	Judger JudgerModels.JudgerConfigType `json:"judger"`
}

type debug struct {
	//Sql bool `json:"sql"`
	Gin bool `json:"gin"`
}

func ReadFile(config *GlobalConfigType) {
	configPath, err := files.ConfigFilePath()
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(content, config); err != nil {
		panic(err)
	}
}
