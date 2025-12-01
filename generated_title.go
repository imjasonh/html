package html

// TitleElement represents the <title> HTML element.

type TitleElement[C HTMLContent] basicElement[C]


// Title creates a new <title> element.
func Title(children ...HTMLContent) *TitleElement[HTMLContent] {
	return &TitleElement[HTMLContent]{
		tagName: "title",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *TitleElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TitleElement[C]) Attr(key, value string) *TitleElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Text sets the text attribute.
func (e *TitleElement[C]) Text(v string) *TitleElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "text", Value: v})
	return e
}


