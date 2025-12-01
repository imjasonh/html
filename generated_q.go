package html

// QuoteElement represents the <q> HTML element.

type QuoteElement[C HTMLContent] basicElement[C]


// Q creates a new <q> element.
func Q(children ...HTMLContent) *QuoteElement[HTMLContent] {
	return &QuoteElement[HTMLContent]{
		tagName: "q",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *QuoteElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *QuoteElement[C]) Attr(key, value string) *QuoteElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Cite sets the cite attribute.
func (e *QuoteElement[C]) Cite(v string) *QuoteElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "cite", Value: v})
	return e
}


