package htmlParser

type HtmlParser interface {
	Create(templateName, fileName string, data interface{}) (string, error)
}
