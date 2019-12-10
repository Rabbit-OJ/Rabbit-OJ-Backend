package config

import (
	"Rabbit-OJ-Backend/utils"
	"encoding/json"
	"io/ioutil"
)

var (
	Global *GlobalConfigType
)

type GlobalConfigType struct {
	Rpc        string     `json:"rpc"`
	RabbitMQ   string     `json:"rabbit_mq"`
	MySQL      string     `json:"mysql"`
	Concurrent concurrent `json:"concurrent"`
	AutoRemove autoRemove `json:"auto_remove"`
}

type concurrent struct {
	Judge uint `json:"judge"`
}
type autoRemove struct {
	Containers bool `json:"containers"`
	Files      bool `json:"files"`
}

func ReadFile(config *GlobalConfigType) {
	configPath, err := utils.ConfigFilePath()
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
