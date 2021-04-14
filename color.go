package color

import (
	"fmt"

	"github.com/fatih/color"
)

var fmtErr = color.New(color.FgRed)

func PrintAppError(app string, msg string) {
	fmtErr.Printf("[%s error]: ", app)
	fmt.Println(msg)
}
