package html

// FrameSetElement represents the <frameset> HTML element.

type FrameSetElement[C HTMLContent] basicElement[C]


// Frameset creates a new <frameset> element.
func Frameset(children ...HTMLContent) *FrameSetElement[HTMLContent] {
	return &FrameSetElement[HTMLContent]{
		tagName: "frameset",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *FrameSetElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *FrameSetElement[C]) Attr(key, value string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Cols sets the cols attribute.
func (e *FrameSetElement[C]) Cols(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "cols", Value: v})
	return e
}



// Onafterprint sets the onafterprint attribute.
func (e *FrameSetElement[C]) Onafterprint(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onafterprint", Value: v})
	return e
}



// Onbeforeprint sets the onbeforeprint attribute.
func (e *FrameSetElement[C]) Onbeforeprint(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onbeforeprint", Value: v})
	return e
}



// Onbeforeunload sets the onbeforeunload attribute.
func (e *FrameSetElement[C]) Onbeforeunload(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onbeforeunload", Value: v})
	return e
}



// Onhashchange sets the onhashchange attribute.
func (e *FrameSetElement[C]) Onhashchange(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onhashchange", Value: v})
	return e
}



// Onlanguagechange sets the onlanguagechange attribute.
func (e *FrameSetElement[C]) Onlanguagechange(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onlanguagechange", Value: v})
	return e
}



// Onmessage sets the onmessage attribute.
func (e *FrameSetElement[C]) Onmessage(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onmessage", Value: v})
	return e
}



// Onmessageerror sets the onmessageerror attribute.
func (e *FrameSetElement[C]) Onmessageerror(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onmessageerror", Value: v})
	return e
}



// Onoffline sets the onoffline attribute.
func (e *FrameSetElement[C]) Onoffline(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onoffline", Value: v})
	return e
}



// Ononline sets the ononline attribute.
func (e *FrameSetElement[C]) Ononline(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "ononline", Value: v})
	return e
}



// Onpagehide sets the onpagehide attribute.
func (e *FrameSetElement[C]) Onpagehide(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpagehide", Value: v})
	return e
}



// Onpagereveal sets the onpagereveal attribute.
func (e *FrameSetElement[C]) Onpagereveal(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpagereveal", Value: v})
	return e
}



// Onpageshow sets the onpageshow attribute.
func (e *FrameSetElement[C]) Onpageshow(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpageshow", Value: v})
	return e
}



// Onpageswap sets the onpageswap attribute.
func (e *FrameSetElement[C]) Onpageswap(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpageswap", Value: v})
	return e
}



// Onpopstate sets the onpopstate attribute.
func (e *FrameSetElement[C]) Onpopstate(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpopstate", Value: v})
	return e
}



// Onrejectionhandled sets the onrejectionhandled attribute.
func (e *FrameSetElement[C]) Onrejectionhandled(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onrejectionhandled", Value: v})
	return e
}



// Onstorage sets the onstorage attribute.
func (e *FrameSetElement[C]) Onstorage(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onstorage", Value: v})
	return e
}



// Onunhandledrejection sets the onunhandledrejection attribute.
func (e *FrameSetElement[C]) Onunhandledrejection(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onunhandledrejection", Value: v})
	return e
}



// Onunload sets the onunload attribute.
func (e *FrameSetElement[C]) Onunload(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onunload", Value: v})
	return e
}



// Rows sets the rows attribute.
func (e *FrameSetElement[C]) Rows(v string) *FrameSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "rows", Value: v})
	return e
}


