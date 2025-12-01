package html

// SpanElement represents the <span> HTML element.

type SpanElement[C HTMLContent] basicElement[C]


// Span creates a new <span> element.
func Span(children ...HTMLContent) *SpanElement[HTMLContent] {
	return &SpanElement[HTMLContent]{
		tagName: "span",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *SpanElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *SpanElement[C]) Attr(key, value string) *SpanElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}

