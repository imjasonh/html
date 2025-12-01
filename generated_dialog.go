package html

// DialogElement represents the <dialog> HTML element.

type DialogElement[C HTMLContent] basicElement[C]


// Dialog creates a new <dialog> element.
func Dialog(children ...HTMLContent) *DialogElement[HTMLContent] {
	return &DialogElement[HTMLContent]{
		tagName: "dialog",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *DialogElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *DialogElement[C]) Attr(key, value string) *DialogElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// ClosedBy sets the closedBy attribute.
func (e *DialogElement[C]) ClosedBy(v string) *DialogElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "closedBy", Value: v})
	return e
}



// Open sets the open attribute.
func (e *DialogElement[C]) Open(v bool) *DialogElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "open", Value: "open"})
	}
	return e
}



// ReturnValue sets the returnValue attribute.
func (e *DialogElement[C]) ReturnValue(v string) *DialogElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "returnValue", Value: v})
	return e
}


