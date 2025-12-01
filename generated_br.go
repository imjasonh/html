package html

// BRElement represents the <br> HTML element.

type BRElement basicElement[HTMLContent]


// Br creates a new <br> element.
func Br() *BRElement {
	return &BRElement{
		tagName: "br",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *BRElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *BRElement) Attr(key, value string) *BRElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Clear sets the clear attribute.
func (e *BRElement) Clear(v string) *BRElement {
	e.attributes = append(e.attributes, Attribute{Key: "clear", Value: v})
	return e
}


