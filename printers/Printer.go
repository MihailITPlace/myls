package printers

import "os"

type Printer interface {
	Print(files []os.FileInfo, printHidden bool)
}
