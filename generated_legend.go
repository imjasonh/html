package html

// LegendElement represents the <legend> HTML element.

type LegendElement[C HTMLContent] basicElement[C]


// Legend creates a new <legend> element.
func Legend(children ...HTMLContent) *LegendElement[HTMLContent] {
	return &LegendElement[HTMLContent]{
		tagName: "legend",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *LegendElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *LegendElement[C]) Attr(key, value string) *LegendElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *LegendElement[C]) Align(v string) *LegendElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}


