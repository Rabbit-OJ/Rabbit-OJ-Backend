package config

import "Rabbit-OJ-Backend/models"

var (
	SupportLanguage []models.SupportLanguage
	CompileObject   map[string]CompileInfo
	LocalImages     map[string]bool
)

type CompileInfo struct {
	BuildArgs   []string    `json:"build_args"`
	Source      string      `json:"source"`
	NoBuild     bool        `json:"no_build"`
	BuildTarget string      `json:"build_target"`
	BuildImage  string      `json:"build_image"`
	Constraints Constraints `json:"constraints"`
	RunArgs     []string    `json:"run_args"`
	RunArgsJSON string      `json:"-"`
	RunImage    string      `json:"run_image"`
}
type Constraints struct {
	BuildTimeout int   `json:"build_timeout"` // unit:seconds
	RunTimeout   int   `json:"run_timeout"`   // unit: seconds
	CPU          int64 `json:"cpu"`           // unit: COREs / 1e9
	Memory       int64 `json:"memory"`        // unit: bytes
}
