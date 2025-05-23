package logger

import (
	"fmt"
	"os"
	"time"

	// col "url-shortener-pro/pkg/color"
	col "github.com/fatih/color"
)

type logger struct {
	ColorInfo      col.Color
	ColorDebug     col.Color
	ColorWarn      col.Color
	ColorError     col.Color
	ColorTime      col.Color
	ColorMessage   col.Color
	ColorArguments col.Color

	CriticalError bool
	Time          bool
	TimeFormat    string
}

type Settings struct {
	ColorInfo      col.Color
	ColorDebug     col.Color
	ColorWarn      col.Color
	ColorError     col.Color
	ColorTime      col.Color
	ColorMessage   col.Color
	ColorArguments col.Color
	CriticalError  bool
	Time           bool
	TimeFormat     string
}

type arg struct {
	Name  string
	Value string
}

func New(settings Settings) logger {
	return logger(settings)
}

func (l *logger) Info(message string, arg ...arg) {
	l.printMessage("INFO", message, l.ColorInfo, arg...)

}

func (l *logger) Debug(message string, arg ...arg) {
	l.printMessage("DEBUG", message, l.ColorDebug, arg...)

}

func (l *logger) Warn(message string, arg ...arg) {
	l.printMessage("WARN", message, l.ColorWarn, arg...)

}

func (l *logger) Error(message string, arg ...arg) {
	l.printMessage("ERROR", message, l.ColorError, arg...)
	if l.CriticalError {
		os.Exit(1)
	}
}

func (l *logger) printMessage(typeMessage string, message string, color col.Color, arg ...arg) {
	if l.Time {
		l.ColorTime.Print(time.Now().Format(l.TimeFormat) + " ")
	}

	color.Print(typeMessage + ": ")

	l.ColorMessage.Print(message)
	if arg != nil {

		l.ColorArguments.Print(" {\n")
		for _, v := range arg {
			fmt.Print("  " + l.ColorArguments.Sprint(v.Name) + ": " + l.ColorArguments.Sprint(v.Value) + " \n")
		}

		l.ColorArguments.Print("}")
	}
	l.ColorArguments.Print("\n")

}

func Arg(name any, value any) arg {
	nameStr := fmt.Sprintf("%v", name)
	valueStr := fmt.Sprintf("%v", value)
	return arg{Name: nameStr, Value: valueStr}
}
