package structs

// GeneralConfig configuration
type GeneralConfig struct {
	FontFamily   string  `json:"font-family"`
	Height       int     `json:"height"`
	Width        int     `json:"width"`
	MarginTop    int     `json:"margin-top"`
	MarginRight  int     `json:"margin-right"`
	MarginLeft   int     `json:"margin-left"`
	MarginBottom int     `json:"margin-bottom"`
	FontSize     int     `json:"font-size"`
	Opacity      float64 `json:"opacity"`
}

// TrayConfig configuration
type TrayConfig struct {
	Padding int `json:"padding"`
}

// VolumeConfig configuration
type VolumeConfig struct {
	Scroll     bool   `json:"scroll"`
	Icon       string `json:"icon"`
	FontWeight string `json:"font-weight"`
	Color      string `json:"color"`
}

// WorkspacesConfig configuration
type WorkspacesConfig struct {
	Click        bool   `json:"click"`
	CurrentColor string `json:"current-color"`
	FontWeight   string `json:"font-weight"`
	Color        string `json:"color"`
}

// Launcher configuration
type Launcher struct {
	Click      bool   `json:"click"`
	Color      string `json:"color"`
	FontWeight string `json:"font-weight"`
}

// PowerConfig configuration
type PowerConfig struct {
	Icon       string `json:"icon"`
	FontWeight string `json:"font-weight"`
	Color      string `json:"color"`
}

// TimeConfig configuration
type TimeConfig struct {
	Click      bool   `json:"click"`
	FontWeight string `json:"font-weight"`
	Color      string `json:"color"`
}

// OlkbConfig configuration
type OlkbConfig struct {
	Enable     bool   `json:"enable"`
	Order      string `json:"order"`
	FontWeight string `json:"font-weight"`
	Color      string `json:"color"`
}

// BarConfig configuration
type BarConfig struct {
	Launcher   Launcher         `json:"launcher"`
	Olkb       OlkbConfig       `json:"olkb"`
	Time       TimeConfig       `json:"time"`
	Tray       TrayConfig       `json:"tray"`
	Power      PowerConfig      `json:"power"`
	Volume     VolumeConfig     `json:"volume"`
	General    GeneralConfig    `json:"general"`
	Workspaces WorkspacesConfig `json:"workspaces"`
}
