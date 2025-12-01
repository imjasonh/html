package html

// HeadElement represents the <head> HTML element.

type HeadElement[C HTMLContent] basicElement[C]


// Head creates a new <head> element.
func Head(children ...HTMLContent) *HeadElement[HTMLContent] {
	return &HeadElement[HTMLContent]{
		tagName: "head",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *HeadElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *HeadElement[C]) Attr(key, value string) *HeadElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}

