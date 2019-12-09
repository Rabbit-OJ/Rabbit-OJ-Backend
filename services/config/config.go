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
}

type concurrent struct {
	JudgeCount uint `json:"judge_count"`
}

func readFile(config *GlobalConfigType) {
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

func Init() {
	Global = &GlobalConfigType{}
	readFile(Global)
}
