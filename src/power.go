package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"./structs"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func updatePower(remaining int, max int, icon string) {
	var builder strings.Builder

	builder.WriteString(icon)
	builder.WriteString("  ")
	builder.WriteString(strconv.Itoa(int(float32(remaining) / float32(max) * 100)))
	builder.WriteByte('%')
	texts["power"].SetText(builder.String())
}

func listenEvents(max int, icon string) {
	var remaining int
	var err error
	var content []byte

	content, err = ioutil.ReadFile("/sys/class/power_supply/BAT1/charge_now")
	if err != nil {
		log.Println(err)
	}
	remaining, err = strconv.Atoi(strings.Trim(string(content), "\n"))
	if err != nil {
		log.Println(err)
	}
	updatePower(remaining, max, icon)
	time.AfterFunc(60000000000, func() { listenEvents(max, icon) })
}

func initPower(config structs.PowerConfig) error {
	var max int
	var err error
	var content []byte

	_, err = os.Stat("/sys/class/power_supply/BAT1/charge_full")
	if os.IsNotExist(err) {
		return nil
	}
	err = nil
	texts["power"] = widgets.NewQLabel(nil, 0)
	texts["power"].SetAlignment(core.Qt__AlignCenter)
	texts["launcher"].SetStyleSheet(fmt.Sprintf("color: white; font-weight: %s", config.FontWeight))
	content, err = ioutil.ReadFile("/sys/class/power_supply/BAT1/charge_full")
	if err != nil {
		return err
	}
	max, err = strconv.Atoi(strings.Trim(string(content), "\n"))
	if err != nil {
		return err
	}
	go listenEvents(max, config.Icon)
	return err
}
