package pdfGenerator

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"os"
)

type wk struct {
	rootPath string
}

func NewWkHtmlToPDF(rootPath string) PDFGeneratorInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile, outputFileName string) (string, error) {
	f, err := os.Open(htmlFile)
	if err != nil {
		return "", err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	outputFilePath := w.rootPath + "/" + outputFileName + ".pdf"
	err = pdfg.WriteFile(outputFilePath)
	if err != nil {
		return "", err
	}

	return outputFilePath, nil
}
