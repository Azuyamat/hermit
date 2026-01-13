package main

import "fmt"

type Color string

const (
	ColorReset  Color = "\033[0m"
	ColorRed    Color = "\033[31m"
	ColorGreen  Color = "\033[32m"
	ColorYellow Color = "\033[33m"
	ColorBlue   Color = "\033[34m"
	ColorPurple Color = "\033[35m"
	ColorCyan   Color = "\033[36m"
	ColorWhite  Color = "\033[37m"
	ColorGray   Color = "\033[90m"
)

func (c Color) String() string {
	return string(c)
}

func Colorize(text string, color Color) string {
	return string(color) + text + string(ColorReset)
}

func (c Color) Sprint(text string) string {
	return Colorize(text, c)
}

func (c Color) Sprintln(text string) string {
	return Colorize(text, c) + "\n"
}

func (c Color) Sprintf(format string, a ...interface{}) string {
	return Colorize(fmt.Sprintf(format, a...), c)
}

func (c Color) Errorf(format string, a ...interface{}) string {
	return Colorize(fmt.Sprintf(format, a...), ColorRed)
}
