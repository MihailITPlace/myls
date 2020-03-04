package printers

import (
	"github.com/fatih/color"
	"os"
)

type BasePrinter struct {
}

func (self *BasePrinter) GetColorName(file os.FileInfo) string {
	clr := color.FgWhite
	if file.IsDir() {
		clr = color.FgYellow
	}
	return color.New(clr).SprintFunc()(file.Name())
}
