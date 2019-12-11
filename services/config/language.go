package config

import "Rabbit-OJ-Backend/models"

var (
	SupportLanguage []models.SupportLanguage
	CompileObject   map[string]CompileInfo
	LocalImages     map[string]bool
)

type CompileInfo struct {
	BuildArgs    []string `json:"build_args"`
	Source       string   `json:"source"`
	NoBuild      bool     `json:"no_build"`
	BuildTarget  string   `json:"build_target"`
	BuildImage   string   `json:"build_image"`
	BuildTimeout int      `json:"build_timeout"` //unit:seconds
	RunArgs      []string `json:"run_args"`
	RunArgsJSON  string   `json:"-"`
	RunImage     string   `json:"run_image"`
	RunTimeout   int      `json:"run_timeout"` // unit: seconds
}
