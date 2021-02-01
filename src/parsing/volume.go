package parsing

import (
	"../structs"
)

func volume(loadedConfig structs.VolumeConfig, config *structs.VolumeConfig) {
	config.Icon = loadedConfig.Icon
	config.Scroll = loadedConfig.Scroll
	config.FontWeight = loadedConfig.FontWeight
	config.Color = loadedConfig.Color
}
