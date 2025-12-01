package html

// ImageElement represents the <img> HTML element.

type ImageElement basicElement[HTMLContent]


// Img creates a new <img> element.
func Img() *ImageElement {
	return &ImageElement{
		tagName: "img",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *ImageElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *ImageElement) Attr(key, value string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *ImageElement) Align(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Alt sets the alt attribute.
func (e *ImageElement) Alt(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "alt", Value: v})
	return e
}



// Complete sets the complete attribute.
func (e *ImageElement) Complete(v bool) *ImageElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "complete", Value: "complete"})
	}
	return e
}



// CurrentSrc sets the currentSrc attribute.
func (e *ImageElement) CurrentSrc(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "currentSrc", Value: v})
	return e
}



// Decoding sets the decoding attribute.
func (e *ImageElement) Decoding(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "decoding", Value: v})
	return e
}



// FetchPriority sets the fetchPriority attribute.
func (e *ImageElement) FetchPriority(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "fetchPriority", Value: v})
	return e
}



// IsMap sets the isMap attribute.
func (e *ImageElement) IsMap(v bool) *ImageElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "isMap", Value: "isMap"})
	}
	return e
}



// Loading sets the loading attribute.
func (e *ImageElement) Loading(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "loading", Value: v})
	return e
}



// LongDesc sets the longDesc attribute.
func (e *ImageElement) LongDesc(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "longDesc", Value: v})
	return e
}



// Lowsrc sets the lowsrc attribute.
func (e *ImageElement) Lowsrc(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "lowsrc", Value: v})
	return e
}



// Name sets the name attribute.
func (e *ImageElement) Name(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// ReferrerPolicy sets the referrerPolicy attribute.
func (e *ImageElement) ReferrerPolicy(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "referrerPolicy", Value: v})
	return e
}



// Sizes sets the sizes attribute.
func (e *ImageElement) Sizes(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "sizes", Value: v})
	return e
}



// Src sets the src attribute.
func (e *ImageElement) Src(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}



// Srcset sets the srcset attribute.
func (e *ImageElement) Srcset(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "srcset", Value: v})
	return e
}



// UseMap sets the useMap attribute.
func (e *ImageElement) UseMap(v string) *ImageElement {
	e.attributes = append(e.attributes, Attribute{Key: "useMap", Value: v})
	return e
}


