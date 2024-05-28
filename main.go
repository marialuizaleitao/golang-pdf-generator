package main

import (
	"fmt"
	"golang-pdf-generator/htmlParser"
)

type Data struct {
	Name string
}

func main() {
	h, err := htmlParser.New("tmp")
	if err != nil {
		fmt.Println("Error creating HTML parser:", err)
		return
	}

	dataHTML := Data{Name: "Maria"}

	fileName, err := h.Create("templates/index.html", "test", dataHTML)
	if err != nil {
		fmt.Println("Error generating HTML file:", err)
		return
	}

	fmt.Println("HTML file generated:", fileName)
}
