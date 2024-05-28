package pdfGenerator

type PDFGeneratorInterface interface {
	Create(htmlFile, fileName string) (string, error)
}
