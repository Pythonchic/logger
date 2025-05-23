package main

import (
	"github.com/Pythonchic/logger"
	color "github.com/fatih/color")


func main () {
	log := logger.New(
		logger.Settings{
			ColorInfo: *color.RGB(42, 110, 219),
			ColorDebug: *color.RGB(247, 255, 0),
			ColorWarn: *color.RGB(255, 119, 0),
			ColorError: *color.RGB(255, 0, 0),
			ColorTime: *color.RGB(150, 232, 56),
			ColorMessage: *color.RGB(48, 200, 227),
			ColorArguments: *color.RGB(224, 224, 224),
			Time: true,
			TimeFormat: "2006/01/02 15:4:5",
		},
	)

	log.Error("Error", logger.Arg("Error", "not found"))
}
