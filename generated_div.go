package html

// DivElement represents the <div> HTML element.

type DivElement[C HTMLContent] basicElement[C]


// Div creates a new <div> element.
func Div(children ...HTMLContent) *DivElement[HTMLContent] {
	return &DivElement[HTMLContent]{
		tagName: "div",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *DivElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *DivElement[C]) Attr(key, value string) *DivElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *DivElement[C]) Align(v string) *DivElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}


