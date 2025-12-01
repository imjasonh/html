package html

// HRElement represents the <hr> HTML element.

type HRElement basicElement[HTMLContent]


// Hr creates a new <hr> element.
func Hr() *HRElement {
	return &HRElement{
		tagName: "hr",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *HRElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *HRElement) Attr(key, value string) *HRElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *HRElement) Align(v string) *HRElement {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Color sets the color attribute.
func (e *HRElement) Color(v string) *HRElement {
	e.attributes = append(e.attributes, Attribute{Key: "color", Value: v})
	return e
}



// NoShade sets the noShade attribute.
func (e *HRElement) NoShade(v bool) *HRElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "noShade", Value: "noShade"})
	}
	return e
}



// Size sets the size attribute.
func (e *HRElement) Size(v string) *HRElement {
	e.attributes = append(e.attributes, Attribute{Key: "size", Value: v})
	return e
}



// Width sets the width attribute.
func (e *HRElement) Width(v string) *HRElement {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: v})
	return e
}


