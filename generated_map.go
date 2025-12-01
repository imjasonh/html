package html

// MapElement represents the <map> HTML element.

type MapElement[C HTMLContent] basicElement[C]


// Map creates a new <map> element.
func Map(children ...HTMLContent) *MapElement[HTMLContent] {
	return &MapElement[HTMLContent]{
		tagName: "map",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *MapElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *MapElement[C]) Attr(key, value string) *MapElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Areas sets the areas attribute.
func (e *MapElement[C]) Areas(v string) *MapElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "areas", Value: v})
	return e
}



// Name sets the name attribute.
func (e *MapElement[C]) Name(v string) *MapElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}


