package html

// FontElement represents the <font> HTML element.

type FontElement[C HTMLContent] basicElement[C]


// Font creates a new <font> element.
func Font(children ...HTMLContent) *FontElement[HTMLContent] {
	return &FontElement[HTMLContent]{
		tagName: "font",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *FontElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *FontElement[C]) Attr(key, value string) *FontElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Face sets the face attribute.
func (e *FontElement[C]) Face(v string) *FontElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "face", Value: v})
	return e
}



// Size sets the size attribute.
func (e *FontElement[C]) Size(v string) *FontElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "size", Value: v})
	return e
}


