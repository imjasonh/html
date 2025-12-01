package html

// HtmlElement represents the <html> HTML element.

type HtmlElement[C HTMLContent] basicElement[C]


// Html creates a new <html> element.
func Html(children ...HTMLContent) *HtmlElement[HTMLContent] {
	return &HtmlElement[HTMLContent]{
		tagName: "html",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *HtmlElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *HtmlElement[C]) Attr(key, value string) *HtmlElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Version sets the version attribute.
func (e *HtmlElement[C]) Version(v string) *HtmlElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "version", Value: v})
	return e
}


