package main

import (
	"fmt"
	"golang-pdf-generator/htmlParser"
)

func main() {
	h, err := htmlParser.New("tmp")
	if err != nil {
		fmt.Println("Error creating HTML parser:", err)
		return
	}

	fileName, err := h.Create("templates/index.html", "test", nil)
	if err != nil {
		fmt.Println("Error generating HTML file:", err)
		return
	}

	fmt.Println("HTML file generated:", fileName)
}
