package html

// BaseElement represents the <base> HTML element.

type BaseElement basicElement[HTMLContent]


// Base creates a new <base> element.
func Base() *BaseElement {
	return &BaseElement{
		tagName: "base",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *BaseElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *BaseElement) Attr(key, value string) *BaseElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Href sets the href attribute.
func (e *BaseElement) Href(v string) *BaseElement {
	e.attributes = append(e.attributes, Attribute{Key: "href", Value: v})
	return e
}



// Target sets the target attribute.
func (e *BaseElement) Target(v string) *BaseElement {
	e.attributes = append(e.attributes, Attribute{Key: "target", Value: v})
	return e
}


