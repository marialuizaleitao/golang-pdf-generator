package pdfGenerator

type PDFGeneratorInterface interface {
	Create(fileName string) (string, error)
}
