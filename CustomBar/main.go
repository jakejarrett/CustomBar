package main

import (
    "fmt"
    "github.com/therecipe/qt/gui"
    "github.com/therecipe/qt/widgets"
)

type BarConfig struct {
    height      int
    width       int
    marginTop   int
    marginRight int
    marginLeft  int
    opacity     float64
    fontSize    int
}

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

var texts    map[string]*widgets.QLabel

func initConfigs(app *widgets.QApplication, config BarConfig) {
    var font    *gui.QFont

    font = gui.NewQFont()
    font.SetPixelSize(config.fontSize)
    app.SetFont(font, "")
}

func main() {
    var err     error
    var app     *widgets.QApplication
    var widget  *widgets.QWidget
    var appName string
    var config  BarConfig

    appName = "custombar"
    texts = make(map[string]*widgets.QLabel)
    err = fillConfig(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    app, widget = initWindow(config)
    initConfigs(app, config)
    initDate()
    err = initPulseAudio(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    createLayout(widget)
    app.Exec()
}
