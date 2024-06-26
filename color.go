package color

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	templateErr     = "[%s error]: "
	templateWarning = "[%s warning]: "
	templateOK      = "[%s ok]: "
)

var (
	fmtErr     = color.New(color.FgRed)
	fmtWarning = color.New(color.FgYellow)
	fmtOK      = color.New(color.FgGreen)
)

func PrintAppError(app string, msg string) {
	printAppColor(app, msg, templateErr, fmtErr)
}

func PrintAppWarning(app string, msg string) {
	printAppColor(app, msg, templateWarning, fmtWarning)
}

func PrintAppOK(app string, msg string) {
	printAppColor(app, msg, templateOK, fmtOK)
}

func printAppColor(app, msg, template string, c *color.Color) {
	c.Printf(template, app)
	fmt.Println(msg)
}
