package html

// SourceElement represents the <source> HTML element.

type SourceElement basicElement[HTMLContent]


// Source creates a new <source> element.
func Source() *SourceElement {
	return &SourceElement{
		tagName: "source",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *SourceElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *SourceElement) Attr(key, value string) *SourceElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Media sets the media attribute.
func (e *SourceElement) Media(v string) *SourceElement {
	e.attributes = append(e.attributes, Attribute{Key: "media", Value: v})
	return e
}



// Sizes sets the sizes attribute.
func (e *SourceElement) Sizes(v string) *SourceElement {
	e.attributes = append(e.attributes, Attribute{Key: "sizes", Value: v})
	return e
}



// Src sets the src attribute.
func (e *SourceElement) Src(v string) *SourceElement {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}



// Srcset sets the srcset attribute.
func (e *SourceElement) Srcset(v string) *SourceElement {
	e.attributes = append(e.attributes, Attribute{Key: "srcset", Value: v})
	return e
}



// Type sets the type attribute.
func (e *SourceElement) Type(v string) *SourceElement {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}


