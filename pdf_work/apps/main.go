package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	pdf.SetTitle("sdfsadfsadf", true)
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println("出库哦呜")
		return
	}
}
