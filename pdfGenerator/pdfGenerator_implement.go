package pdfGenerator

type wkhtml struct {
	rootPath string
}

func New(rootPath string) PDFGeneratorInterface {
	return &wkhtml{rootPath: rootPath}
}

func (w *wkhtml) Create(htmlFile string) (string, error) {
	return "", nil
}
