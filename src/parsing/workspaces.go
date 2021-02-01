package parsing

import (
	"../structs"
)

func workspaces(loadedConfig structs.WorkspacesConfig, config *structs.WorkspacesConfig) {
	config.Click = loadedConfig.Click
	config.CurrentColor = loadedConfig.CurrentColor
	config.FontWeight = loadedConfig.FontWeight
	config.Color = loadedConfig.Color
}
