package initialize

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/services/config"
	"os"
)

func Config() {
	config.Global = &config.GlobalConfigType{}
	config.ReadFile(config.Global)
	config.Secret = os.Getenv("Secret")
	Language()
}

func Language() {
	languageCount := len(config.Global.Languages)

	config.LocalImages = map[string]bool{}
	config.CompileObject = map[string]*config.CompileInfo{}
	config.SupportLanguage = make([]models.SupportLanguage, languageCount)

	for _, item := range config.Global.LocalImages {
		config.LocalImages[item] = true
	}
	for index, item := range config.Global.Languages {
		config.SupportLanguage[index] = models.SupportLanguage{
			Name:  item.Name,
			Value: item.ID,
		}

		config.CompileObject[item.ID] = &item.Args
	}
}
