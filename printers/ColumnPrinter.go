package printers

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type ColumnPrinter struct {
	BasePrinter
}

func NewColumnPrinter() *ColumnPrinter {
	return new(ColumnPrinter)
}

func (self *ColumnPrinter) Print(files []os.FileInfo, printHidden bool) {
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)

	str := ""
	i := 0
	cnt := 0
	for {
		if i >= len(files) {
			break
		}

		if !printHidden && files[i].Name()[0] == '.' {
			i++
			continue
		}

		str += self.BasePrinter.GetColorName(files[i]) + "\t"

		cnt++
		i++

		if cnt%3 == 0 {
			fmt.Fprintln(w, str)
			str = ""
			continue
		}
	}

	if str != "" {
		fmt.Fprintln(w, str)
	}

	w.Flush()
}
