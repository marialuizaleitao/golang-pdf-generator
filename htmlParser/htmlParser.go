package htmlParser

type HtmlParserInterface interface {
	Create(templateName, fileName string, data interface{}) (string, error)
}
