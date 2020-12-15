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
	config.FontFamily = loadedConfig.FontFamily

	if width == 0 {
		width = defaultWidth
	}

	config.Width = width
}
