package html

// TableCaptionElement represents the <caption> HTML element.

type TableCaptionElement[C HTMLContent] basicElement[C]


// Caption creates a new <caption> element.
func Caption(children ...HTMLContent) *TableCaptionElement[HTMLContent] {
	return &TableCaptionElement[HTMLContent]{
		tagName: "caption",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *TableCaptionElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TableCaptionElement[C]) Attr(key, value string) *TableCaptionElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *TableCaptionElement[C]) Align(v string) *TableCaptionElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}


