package config

import "Rabbit-OJ-Backend/models"

var (
	SupportLanguage []models.SupportLanguage
	CompileObject   map[string]*CompileInfo
	LocalImages     map[string]bool
)

type CompileInfo struct {
	BuildArgs     []string
	Source        string
	NoBuild       bool //todo: support no-build language like: nodejs & python ...
	BuildTarget   string
	BuildImage    string
	BuildTimeout  int //unit:seconds
	RunArgs       string
	RunImage      string
	RunMaxTimeout int // unit: seconds
}
