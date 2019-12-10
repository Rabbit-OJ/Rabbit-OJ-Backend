package initialize

import "Rabbit-OJ-Backend/services/config"

func Config() {
	config.Global = &config.GlobalConfigType{}
	config.ReadFile(config.Global)
}
