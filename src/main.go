package main

// #include "./events.h"
// #include "./tray.h"
// #cgo pkg-config: x11 xcb xcb-util
import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"./parsing"
	"./structs"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func errorHandler(err error) {
	fmt.Fprintf(os.Stderr, "An error occured: %v\n", err)
}

var texts map[string]*widgets.QLabel

func initConfigs(app *widgets.QApplication, config structs.BarConfig) {
	var font *gui.QFont

	font = gui.NewQFont2(config.General.FontFamily, 11, int(gui.QFont__Normal), false)
	font.SetPixelSize(config.General.FontSize)
	app.SetFont(font, "")
}

func main() {
	var err error
	var appName string
	var signals *Signals
	var config structs.BarConfig
	var widget *widgets.QWidget
	var app *widgets.QApplication

	appName = "custombar"
	texts = make(map[string]*widgets.QLabel)
	if err != nil {
		errorHandler(err)
		return
	}
	app = widgets.NewQApplication(len(os.Args), os.Args)
	widget = widgets.NewQWidget(nil, core.Qt__FramelessWindowHint)
	screen := app.Desktop().ScreenGeometry(1)
	err = parsing.FillConfig(appName, &config, screen.Width())
	if err != nil {
		errorHandler(err)
		return
	}
	initWindow(config.General, widget, screen)
	initConfigs(app, config)
	initLauncher(config.Launcher, config.General.Height, app)
	if err != nil {
		errorHandler(err)
		return
	}
	err = initPower(config.Power)
	if err != nil {
		errorHandler(err)
		return
	}
	signals = NewSignals(nil)
	err = initPulseAudio(appName, unsafe.Pointer(signals), config.Volume)
	if err != nil {
		errorHandler(err)
		return
	}
	initDate(signals, config.Time)
	initOlkb(signals, config.Olkb)
	createLayout(widget, config.General)
	go C.createTrayManager(C.ulong(config.General.Width), C.ulong(config.General.Height), C.ulong(config.General.Opacity), C.ulong(config.Tray.Padding), unsafe.Pointer(widget.Layout().ItemAt(2).Layout()))
	app.Exec()
}
