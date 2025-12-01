package html

// EmbedElement represents the <embed> HTML element.

type EmbedElement basicElement[HTMLContent]


// Embed creates a new <embed> element.
func Embed() *EmbedElement {
	return &EmbedElement{
		tagName: "embed",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *EmbedElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *EmbedElement) Attr(key, value string) *EmbedElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *EmbedElement) Align(v string) *EmbedElement {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Height sets the height attribute.
func (e *EmbedElement) Height(v string) *EmbedElement {
	e.attributes = append(e.attributes, Attribute{Key: "height", Value: v})
	return e
}



// Name sets the name attribute.
func (e *EmbedElement) Name(v string) *EmbedElement {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Src sets the src attribute.
func (e *EmbedElement) Src(v string) *EmbedElement {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}



// Type sets the type attribute.
func (e *EmbedElement) Type(v string) *EmbedElement {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// Width sets the width attribute.
func (e *EmbedElement) Width(v string) *EmbedElement {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: v})
	return e
}


