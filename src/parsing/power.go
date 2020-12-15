package parsing

import (
	"../structs"
)

func power(loadedConfig structs.PowerConfig, config *structs.PowerConfig) {
	config.Icon = loadedConfig.Icon
}
