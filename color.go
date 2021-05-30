package color

import (
	"fmt"

	"github.com/fatih/color"
)

const templateErr = "[%s error]: "

var fmtErr = color.New(color.FgRed)

func PrintAppError(app string, msg string) {
	printAppColor(app, msg, templateErr, fmtErr)
}

func printAppColor(app, msg, template string, c *color.Color) {
	c.Printf(template, app)
	fmt.Println(msg)
}
