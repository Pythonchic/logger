package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"
)

type logger struct {
	ColorInfo      color.RGBColor
	ColorDebug     color.RGBColor
	ColorWarn      color.RGBColor
	ColorError     color.RGBColor
	ColorMessage   color.RGBColor
	ColorTime      color.RGBColor
	ColorArguments color.RGBColor

	CriticalError bool

	Time       bool
	TimeFormat string

	PrettyArguments bool
}

type arg struct {
	Name  string
	Value string
}

func NewLogger() logger {
	return logger{
		ColorInfo:      color.RGB(45, 157, 237),
		ColorDebug:     color.RGB(189, 0, 142),
		ColorWarn:      color.RGB(255, 119, 0),
		ColorError:     color.RGB(255, 0, 0),
		ColorTime:      color.RGB(150, 232, 56),
		ColorMessage:   color.RGB(48, 200, 227),
		ColorArguments: color.RGB(247, 247, 247),

		CriticalError: false,

		Time:       true,
		TimeFormat: "2006/01/02 15:4:5",

		PrettyArguments: false,
	}
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

func (l *logger) printMessage(typeMessage string, message string, color color.RGBColor, arg ...arg) {
	if l.Time {
		l.ColorTime.Print(time.Now().Format(l.TimeFormat) + " ")
	}

	color.Print(typeMessage + ": ")

	l.ColorMessage.Print(message)
	if arg != nil {
		l.ColorArguments.Print(" {")
		if l.PrettyArguments {
			l.ColorArguments.Print("\n")
			for _, v := range arg {
				fmt.Print(l.ColorArguments.Sprint("  " + v.Name + ": " + v.Value + " \n"))
			}
		} else {
			fmt.Print(l.ColorArguments.Sprint(" " + arg[0].Name + ": " + arg[0].Value))

			if len(arg) > 1 {
				for _, v := range arg[1:] {
					fmt.Print(l.ColorArguments.Sprint(", " + v.Name + ": " + v.Value))
				}
			} else {
				fmt.Print(l.ColorArguments.Sprint(" "))
			}
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
