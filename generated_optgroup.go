package html

// OptGroupElement represents the <optgroup> HTML element.

type OptGroupElement[C HTMLContent] basicElement[C]


// Optgroup creates a new <optgroup> element.
func Optgroup(children ...HTMLContent) *OptGroupElement[HTMLContent] {
	return &OptGroupElement[HTMLContent]{
		tagName: "optgroup",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *OptGroupElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *OptGroupElement[C]) Attr(key, value string) *OptGroupElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Disabled sets the disabled attribute.
func (e *OptGroupElement[C]) Disabled(v bool) *OptGroupElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// Label sets the label attribute.
func (e *OptGroupElement[C]) Label(v string) *OptGroupElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "label", Value: v})
	return e
}


