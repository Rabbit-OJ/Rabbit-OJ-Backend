package config

import (
	"Rabbit-OJ-Backend/models"
	JuderModels "Rabbit-OJ-Backend/services/judger/models"
)

var (
	SupportLanguage []models.SupportLanguage
	CompileObject   map[string]JuderModels.CompileInfo
	LocalImages     map[string]bool
)
