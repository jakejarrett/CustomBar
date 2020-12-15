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

func createDefaultConfig() structs.BarConfig {
	config := structs.BarConfig{
		General: structs.GeneralConfig{
			FontFamily:   "Roboto",
			Height:       33,
			Width:        0,
			Opacity:      40,
			FontSize:     16,
			MarginTop:    0,
			MarginRight:  0,
			MarginLeft:   0,
			MarginBottom: 0,
		},
		Launcher: structs.Launcher{
			Click: true,
			Color: "white",
		},
		Time: structs.TimeConfig{
			Click: true,
		},
		Tray: structs.TrayConfig{
			Padding: 5,
		},
		Power: structs.PowerConfig{
			Icon: "",
		},
		Volume: structs.VolumeConfig{
			Icon:   "",
			Scroll: true,
		},
		Workspaces: structs.WorkspacesConfig{
			Click:        true,
			CurrentColor: "#0053a0",
		},
	}

	return config
}

func getJSONFile(path string) structs.BarConfig {
	var config structs.BarConfig
	file := fmt.Sprintf("%s/config.json", path)

	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)

		os.Mkdir(path, 0777)
		f, fErr := os.Create(file)

		if fErr != nil {
			fmt.Println(fErr)
		}

		defaultConfig := createDefaultConfig()

		defaultConf, defaultConfErr := json.Marshal(createDefaultConfig())

		if defaultConfErr != nil {
			fmt.Println(defaultConfErr)
		}

		f.Write(defaultConf)

		return defaultConfig
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config)

	return config
}

// FillConfig will read the config or create a default config.
func FillConfig(appName string, config *structs.BarConfig, width int) error {
	jsonPath := fmt.Sprintf("%s/.config/%s", os.Getenv("HOME"), appName)
	jsonContent := getJSONFile(jsonPath)
	handleJSONConfig(jsonContent, config, width)
	return nil
}
