package html

// MetaElement represents the <meta> HTML element.

type MetaElement basicElement[HTMLContent]


// Meta creates a new <meta> element.
func Meta() *MetaElement {
	return &MetaElement{
		tagName: "meta",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *MetaElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *MetaElement) Attr(key, value string) *MetaElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Content sets the content attribute.
func (e *MetaElement) Content(v string) *MetaElement {
	e.attributes = append(e.attributes, Attribute{Key: "content", Value: v})
	return e
}



// HttpEquiv sets the httpEquiv attribute.
func (e *MetaElement) HttpEquiv(v string) *MetaElement {
	e.attributes = append(e.attributes, Attribute{Key: "httpEquiv", Value: v})
	return e
}



// Media sets the media attribute.
func (e *MetaElement) Media(v string) *MetaElement {
	e.attributes = append(e.attributes, Attribute{Key: "media", Value: v})
	return e
}



// Name sets the name attribute.
func (e *MetaElement) Name(v string) *MetaElement {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Scheme sets the scheme attribute.
func (e *MetaElement) Scheme(v string) *MetaElement {
	e.attributes = append(e.attributes, Attribute{Key: "scheme", Value: v})
	return e
}


