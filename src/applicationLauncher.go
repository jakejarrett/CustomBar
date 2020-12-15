package main

import (
	"fmt"

	"./structs"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func initLauncher(config structs.Launcher) {
	var filter *core.QObject

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
				fmt.Printf("Launcher should open\n")
			}
			return false
		})
		texts["launcher"].InstallEventFilter(filter)
	}
}
