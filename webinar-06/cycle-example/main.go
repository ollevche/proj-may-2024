package main

import (
	"cycle/paper"
	"cycle/printer"
	"fmt"
)

type customStorer int

func (s customStorer) Store(_ paper.Paper) {
	fmt.Println("Custom storer")
}

func main() {
	printer := &printer.Printer{}

	printer.Print()

	var paper paper.Paper = "Hello world!"

	paper.StoreIn(printer)

	printer.Print()

	checkPrinter(printer)

	checkPrinter(customStorer(1))
}

func checkPrinter(p paper.Storer) {
	// printerP := p.(*printer.Printer)
	printerP, ok := p.(*printer.Printer) // DANGER! No compile time checks
	if !ok {
		fmt.Println("Not a printer")
		return
	}

	printerP.Print()
}
