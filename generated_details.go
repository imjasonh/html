package html

// DetailsElement represents the <details> HTML element.

type DetailsElement[C HTMLContent] basicElement[C]


// Details creates a new <details> element.
func Details(children ...HTMLContent) *DetailsElement[HTMLContent] {
	return &DetailsElement[HTMLContent]{
		tagName: "details",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *DetailsElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *DetailsElement[C]) Attr(key, value string) *DetailsElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Name sets the name attribute.
func (e *DetailsElement[C]) Name(v string) *DetailsElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Open sets the open attribute.
func (e *DetailsElement[C]) Open(v bool) *DetailsElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "open", Value: "open"})
	}
	return e
}


