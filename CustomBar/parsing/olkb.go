package parsing

import (
	"../structs"
)

func olkb(loadedConfig structs.OlkbConfig, config *structs.OlkbConfig) {
	config.Enable = loadedConfig.Enable
	config.Order = loadedConfig.Order
}
