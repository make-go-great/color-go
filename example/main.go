package main

import (
	"github.com/make-go-great/color-go"
)

func main() {
	color.PrintAppError("app", "error message")
	color.PrintAppWarning("app", "warning message")
	color.PrintAppOK("app", "ok message")
}
