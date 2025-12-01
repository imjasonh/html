package html

// LinkElement represents the <link> HTML element.

type LinkElement basicElement[HTMLContent]


// Link creates a new <link> element.
func Link() *LinkElement {
	return &LinkElement{
		tagName: "link",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *LinkElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *LinkElement) Attr(key, value string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// As sets the as attribute.
func (e *LinkElement) As(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "as", Value: v})
	return e
}



// Blocking sets the blocking attribute.
func (e *LinkElement) Blocking(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "blocking", Value: v})
	return e
}



// Charset sets the charset attribute.
func (e *LinkElement) Charset(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "charset", Value: v})
	return e
}



// Disabled sets the disabled attribute.
func (e *LinkElement) Disabled(v bool) *LinkElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// FetchPriority sets the fetchPriority attribute.
func (e *LinkElement) FetchPriority(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "fetchPriority", Value: v})
	return e
}



// Href sets the href attribute.
func (e *LinkElement) Href(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "href", Value: v})
	return e
}



// Hreflang sets the hreflang attribute.
func (e *LinkElement) Hreflang(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "hreflang", Value: v})
	return e
}



// ImageSizes sets the imageSizes attribute.
func (e *LinkElement) ImageSizes(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "imageSizes", Value: v})
	return e
}



// ImageSrcset sets the imageSrcset attribute.
func (e *LinkElement) ImageSrcset(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "imageSrcset", Value: v})
	return e
}



// Integrity sets the integrity attribute.
func (e *LinkElement) Integrity(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "integrity", Value: v})
	return e
}



// Media sets the media attribute.
func (e *LinkElement) Media(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "media", Value: v})
	return e
}



// ReferrerPolicy sets the referrerPolicy attribute.
func (e *LinkElement) ReferrerPolicy(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "referrerPolicy", Value: v})
	return e
}



// Rel sets the rel attribute.
func (e *LinkElement) Rel(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "rel", Value: v})
	return e
}



// RelList sets the relList attribute.
func (e *LinkElement) RelList(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "relList", Value: v})
	return e
}



// Rev sets the rev attribute.
func (e *LinkElement) Rev(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "rev", Value: v})
	return e
}



// Sizes sets the sizes attribute.
func (e *LinkElement) Sizes(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "sizes", Value: v})
	return e
}



// Target sets the target attribute.
func (e *LinkElement) Target(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "target", Value: v})
	return e
}



// Type sets the type attribute.
func (e *LinkElement) Type(v string) *LinkElement {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}


