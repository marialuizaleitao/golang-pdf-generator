package main

import (
	"fmt"
	"golang-pdf-generator/htmlParser"
	"golang-pdf-generator/pdfGenerator"
	"os"
)

type Data struct {
	Name string
}

func main() {
	wk := pdfGenerator.NewWkHtmlToPDF("tmp")
	h, err := htmlParser.New("tmp")
	if err != nil {
		fmt.Println("Error creating HTML parser:", err)
		return
	}

	dataHTML := Data{Name: "Maria"}

	htmlGenerated, err := h.Create("templates/index.html", "test", dataHTML)
	if err != nil {
		fmt.Println("Error generating HTML file:", err)
		return
	}

	fmt.Println("HTML file generated:", htmlGenerated)
	defer os.Remove(htmlGenerated)

	pdfGenerated, err := wk.Create(htmlGenerated, "test")
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return
	}

	fmt.Println("PDF generated:", pdfGenerated)
}
