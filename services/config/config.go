package config

import (
	"Rabbit-OJ-Backend/utils/files"
	"encoding/json"
	"io/ioutil"
)

var (
	Global *GlobalConfigType
)

type GlobalConfigType struct {
	Rpc         string     `json:"rpc"`
	RabbitMQ    string     `json:"rabbit_mq"`
	MySQL       string     `json:"mysql"`
	AutoRemove  autoRemove `json:"auto_remove"`
	Concurrent  concurrent `json:"concurrent"`
	LocalImages []string   `json:"local_images"`
	Languages   []language `json:"languages"`
	Extensions  extensions `json:"extensions"`
}

type extensions struct {
	Dind       bool       `json:"dind"`
	AutoPull   bool       `json:"auto_pull"`
	CheckJudge checkJudge `json:"check_judge"`
	Expire     expire     `json:"expire"`
	Debug      debug      `json:"debug"`
}
type debug struct {
	Sql bool `json:"sql"`
	Gin bool `json:"gin"`
}
type expire struct {
	Enabled  bool  `json:"enabled"`
	Interval int64 `json:"interval"` // interval: minutes
}
type checkJudge struct {
	Enabled  bool  `json:"enabled"`
	Interval int64 `json:"interval"` // interval: minutes
	Requeue  bool  `json:"requeue"`
}
type language struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Enabled bool        `json:"enabled"`
	Args    CompileInfo `json:"args"`
}
type concurrent struct {
	Judge uint `json:"judge"`
}
type autoRemove struct {
	Containers bool `json:"containers"`
	Files      bool `json:"files"`
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
