package parsing

import (
	"../structs"
)

func tray(loadedConfig structs.TrayConfig, config *structs.TrayConfig) {
	config.Padding = loadedConfig.Padding
}
