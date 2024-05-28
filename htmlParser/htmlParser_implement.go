package htmlParser

import (
	"fmt"
	"os"
	"text/template"
)

type htmlStruct struct {
	rootPath string
}

func New(rootPath string) (HtmlParserInterface, error) {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		if err := os.MkdirAll(rootPath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	}

	return &htmlStruct{rootPath: rootPath}, nil
}

func (h *htmlStruct) Create(templateName, name string, data interface{}) (string, error) {
	templateGenerator, err := template.ParseFiles(templateName)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	fileName := h.rootPath + "/" + name + ".html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer fileWriter.Close() // Ensure file is closed after writing

	if err := templateGenerator.Execute(fileWriter, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return fileName, nil
}
