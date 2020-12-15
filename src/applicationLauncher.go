package main

import (
	"fmt"

	"./structs"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func initLauncher(config structs.Launcher, heightOffset int, app *widgets.QApplication) {
	launcherShow := false
	var filter *core.QObject

	launcher := widgets.NewQWidget(nil, core.Qt__FramelessWindowHint)
	launcher.SetFixedHeight(900)
	launcher.SetFixedWidth(600)
	launcher.Move2(0, heightOffset)
	launcher.SetAttribute(core.Qt__WA_X11NetWmWindowTypeUtility, true)
	launcher.SetAttribute(core.Qt__WA_X11NetWmWindowTypeSplash, true)
	launcher.SetStyleSheet("background-color: rgba(0, 0, 0, 102)")

	launcher.ConnectChangeEvent(func(event *core.QEvent) {
		isActive := launcher.IsActiveWindow()
		if !isActive {
			launcherShow = false
			launcher.Hide()
		}
	})

	texts["launcher"] = widgets.NewQLabel(nil, 0)
	texts["launcher"].SetText("Launcher")
	texts["launcher"].SetMinimumWidth(40)
	texts["launcher"].SetAlignment(core.Qt__AlignHCenter | core.Qt__AlignVCenter)
	texts["launcher"].SetStyleSheet(fmt.Sprintf("color: %s", config.Color))
	texts["launcher"].SetEnabled(true)
	if config.Click {
		filter = core.NewQObject(nil)
		filter.ConnectEventFilter(func(watched *core.QObject, event *core.QEvent) bool {
			if event.Type() == core.QEvent__MouseButtonPress {
				if launcherShow {
					launcher.Hide()
				} else {
					launcher.Show()
					app.SetActiveWindow(launcher)
				}

				launcherShow = !launcherShow
			}
			return false
		})
		texts["launcher"].InstallEventFilter(filter)
	}
}
