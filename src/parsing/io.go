package parsing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../structs"
)

func handleJSONConfig(loadedConfig structs.BarConfig, config *structs.BarConfig, defaultWidth int) {

	general(loadedConfig.General, &config.General, defaultWidth)
	power(loadedConfig.Power, &config.Power)
	workspaces(loadedConfig.Workspaces, &config.Workspaces)
	launcher(loadedConfig.Launcher, &config.Launcher)
	tray(loadedConfig.Tray, &config.Tray)
	volume(loadedConfig.Volume, &config.Volume)
	time(loadedConfig.Time, &config.Time)
	olkb(loadedConfig.Olkb, &config.Olkb)
}

func defaultConfig(config *structs.BarConfig, width int) {
	config.General.Height = 33
	config.General.Height = 33
	config.General.Width = width
	config.General.MarginTop = 0
	config.General.MarginLeft = 0
	config.General.MarginRight = 0
	config.General.Opacity = 40
	config.General.FontSize = 16
	config.Workspaces.CurrentColor = "#0053a0"
	config.Workspaces.Click = true
	config.Volume.Icon = ""
	config.Volume.Scroll = true
	config.Power.Icon = ""
	config.Tray.Padding = 5
	config.Time.Click = true
	config.Olkb.Enable = false
	config.Olkb.Order = ""
	config.Launcher.Click = false
	config.Launcher.Color = "white"
}

func getJSONFile(path string) structs.BarConfig {
	var config structs.BarConfig
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)

	return config
}

// FillConfig will read the config or create a default config.
func FillConfig(appName string, config *structs.BarConfig, width int) error {
	defaultConfig(config, width)
	jsonPath := fmt.Sprintf("%s/.config/%s/config.json", os.Getenv("HOME"), appName)
	jsonContent := getJSONFile(jsonPath)
	handleJSONConfig(jsonContent, config, width)
	return nil
}
