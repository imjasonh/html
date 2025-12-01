package html

// UListElement represents the <ul> HTML element.

type UListElement[C HTMLContent] basicElement[C]


// Ul creates a new <ul> element.
func Ul(children ...HTMLContent) *UListElement[HTMLContent] {
	return &UListElement[HTMLContent]{
		tagName: "ul",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *UListElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *UListElement[C]) Attr(key, value string) *UListElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Compact sets the compact attribute.
func (e *UListElement[C]) Compact(v bool) *UListElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "compact", Value: "compact"})
	}
	return e
}



// Type sets the type attribute.
func (e *UListElement[C]) Type(v string) *UListElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}


