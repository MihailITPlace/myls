package printers

import (
	"fmt"
	"os"
)

type SimplePrinter struct {
	BasePrinter
}

func NewSimplePrinter() *SimplePrinter {
	return new(SimplePrinter)
}

func (self *SimplePrinter) Print(files []os.FileInfo, printHidden bool) {
	for _, i := range files {
		if !printHidden && i.Name()[0] == '.' {
			continue
		}
		fmt.Println(self.GetColorName(i))
	}
}
