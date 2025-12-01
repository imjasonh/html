package html

// UnknownElement represents the <unknown> HTML element.

type UnknownElement[C HTMLContent] basicElement[C]


// Unknown creates a new <unknown> element.
func Unknown(children ...HTMLContent) *UnknownElement[HTMLContent] {
	return &UnknownElement[HTMLContent]{
		tagName: "unknown",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *UnknownElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *UnknownElement[C]) Attr(key, value string) *UnknownElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}

