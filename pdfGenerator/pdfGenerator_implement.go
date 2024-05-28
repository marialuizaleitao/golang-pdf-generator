package pdfGenerator

type wk struct {
	rootPath string
}

func NewWkHtmlToPDF(rootPath string) PDFGeneratorInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile string) (string, error) {
	return "", nil
}
