package main

import (
	"../logic"
	"../printers"
	"flag"
)

func main() {

	printHidden := flag.Bool("a", false, "Print all files")
	useWideFormat := flag.Bool("l", false, "Use wide format")
	humanReadable := flag.Bool("h", false, "Human readable. Use with -l")
	useColumnFormat := flag.Bool("C", false, "Use column format")

	flag.Parse()

	files := logic.GetAllFiles()
	var printer printers.Printer

	if *useWideFormat {
		printer = printers.NewWideFormatPrinter(*humanReadable)
	} else if *useColumnFormat {
		printer = printers.NewColumnPrinter()
	} else {
		printer = printers.NewSimplePrinter()
	}

	printer.Print(files, *printHidden)
}
