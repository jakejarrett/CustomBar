package parsing

import (
	"../structs"
)

func launcher(loadedConfig structs.Launcher, config *structs.Launcher) {
	config.Click = loadedConfig.Click
	config.Color = loadedConfig.Color
}
