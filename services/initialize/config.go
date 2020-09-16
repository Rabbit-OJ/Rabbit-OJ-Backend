package initialize

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	JudgerModels "Rabbit-OJ-Backend/services/judger/models"
	"encoding/json"
	"os"
)

func Config() {
	config.Global = &config.GlobalConfigType{}
	config.ReadFile(config.Global)
	config.Secret = os.Getenv("Secret")
	Language()
}

func Language() {
	languageCount := 0
	for _, item := range config.Global.Judger.Languages {
		if item.Enabled {
			languageCount++
		}
	}

	config.LocalImages = map[string]bool{}
	config.CompileObject = map[string]JudgerModels.CompileInfo{}
	config.SupportLanguage = make([]models.SupportLanguage, languageCount)

	for _, item := range config.Global.Judger.LocalImages {
		config.LocalImages[item] = true
	}
	for index, item := range config.Global.Judger.Languages {
		if !item.Enabled {
			continue
		}

		config.SupportLanguage[index] = models.SupportLanguage{
			Name:  item.Name,
			Value: item.ID,
		}

		runArgsJson, err := json.Marshal(item.Args.RunArgs)
		if err != nil {
			panic(err)
		}

		currentCompileObject := item.Args
		currentCompileObject.RunArgsJSON = string(runArgsJson)
		config.CompileObject[item.ID] = currentCompileObject
	}
}
