package html

// DListElement represents the <dl> HTML element.

type DListElement[C HTMLContent] basicElement[C]


// Dl creates a new <dl> element.
func Dl(children ...HTMLContent) *DListElement[HTMLContent] {
	return &DListElement[HTMLContent]{
		tagName: "dl",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *DListElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *DListElement[C]) Attr(key, value string) *DListElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Compact sets the compact attribute.
func (e *DListElement[C]) Compact(v bool) *DListElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "compact", Value: "compact"})
	}
	return e
}


