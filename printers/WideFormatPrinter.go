package printers

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type WideFormatPrinter struct {
	BasePrinter
	_humanReadable bool
}

func NewWideFormatPrinter(humanReadable bool) *WideFormatPrinter {
	printer := new(WideFormatPrinter)
	printer._humanReadable = humanReadable
	return printer
}

func (self *WideFormatPrinter) Print(files []os.FileInfo, printHidden bool) {
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Режим\tРазмер\tИмя\tВремя изменения\t")

	for _, i := range files {
		if !printHidden && i.Name()[0] == '.' {
			continue
		}

		var size string
		if !self._humanReadable {
			size = strconv.FormatInt(i.Size(), 10)
		} else {
			size = getReadableSize(i.Size())
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", i.Mode().String(), size, self.GetColorName(i), i.ModTime().String())
	}

	w.Flush()
}

const BYTE = 1
const KILOBYTE = 1024
const MEGABYTE = 1048576
const GIGABYTE = 1073741824

func getReadableSize(size int64) string {
	if size > GIGABYTE {
		return calcSize(size, GIGABYTE, 'G')
	}

	if size > MEGABYTE {
		return calcSize(size, MEGABYTE, 'M')
	}

	if size > KILOBYTE {
		return calcSize(size, KILOBYTE, 'K')
	}

	return calcSize(size, BYTE, 'B')
}

func calcSize(size int64, divider int64, char int32) string {
	res := float64(size) / float64(divider)
	return fmt.Sprintf("%.1f%c", res, char)
}
