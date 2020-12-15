package parsing

import (
	"../structs"
)

func general(loadedConfig structs.GeneralConfig, config *structs.GeneralConfig, defaultWidth int) {
	width := loadedConfig.Width
	config.MarginTop = loadedConfig.MarginTop
	config.MarginRight = loadedConfig.MarginRight
	config.MarginLeft = loadedConfig.MarginLeft
	config.Height = loadedConfig.Height
	config.Opacity = loadedConfig.Opacity
	config.FontSize = loadedConfig.FontSize

	if width == 0 {
		width = defaultWidth
	}

	config.Width = width
}
