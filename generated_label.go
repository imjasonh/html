package html

// LabelElement represents the <label> HTML element.

type LabelElement[C HTMLContent] basicElement[C]


// Label creates a new <label> element.
func Label(children ...HTMLContent) *LabelElement[HTMLContent] {
	return &LabelElement[HTMLContent]{
		tagName: "label",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *LabelElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *LabelElement[C]) Attr(key, value string) *LabelElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// HtmlFor sets the htmlFor attribute.
func (e *LabelElement[C]) HtmlFor(v string) *LabelElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "htmlFor", Value: v})
	return e
}


