package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"os"
)

func Config() {
	config.Global = &config.GlobalConfigType{}
	config.ReadFile(config.Global)
	config.Secret = os.Getenv("Secret")
}
