package html

// ParagraphElement represents the <p> HTML element.

type ParagraphElement[C HTMLContent] basicElement[C]


// P creates a new <p> element.
func P(children ...HTMLContent) *ParagraphElement[HTMLContent] {
	return &ParagraphElement[HTMLContent]{
		tagName: "p",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *ParagraphElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *ParagraphElement[C]) Attr(key, value string) *ParagraphElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *ParagraphElement[C]) Align(v string) *ParagraphElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}


