package parsing

import (
	"../structs"
)

func time(loadedConfig structs.TimeConfig, config *structs.TimeConfig) {
	config.Click = loadedConfig.Click
}
