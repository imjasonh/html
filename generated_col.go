package html

// TableColElement represents the <col> HTML element.

type TableColElement basicElement[HTMLContent]


// Col creates a new <col> element.
func Col() *TableColElement {
	return &TableColElement{
		tagName: "col",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *TableColElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TableColElement) Attr(key, value string) *TableColElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *TableColElement) Align(v string) *TableColElement {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Ch sets the ch attribute.
func (e *TableColElement) Ch(v string) *TableColElement {
	e.attributes = append(e.attributes, Attribute{Key: "ch", Value: v})
	return e
}



// ChOff sets the chOff attribute.
func (e *TableColElement) ChOff(v string) *TableColElement {
	e.attributes = append(e.attributes, Attribute{Key: "chOff", Value: v})
	return e
}



// VAlign sets the vAlign attribute.
func (e *TableColElement) VAlign(v string) *TableColElement {
	e.attributes = append(e.attributes, Attribute{Key: "vAlign", Value: v})
	return e
}



// Width sets the width attribute.
func (e *TableColElement) Width(v string) *TableColElement {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: v})
	return e
}


