package html

// BodyElement represents the <body> HTML element.

type BodyElement[C HTMLContent] basicElement[C]


// Body creates a new <body> element.
func Body(children ...HTMLContent) *BodyElement[HTMLContent] {
	return &BodyElement[HTMLContent]{
		tagName: "body",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *BodyElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *BodyElement[C]) Attr(key, value string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Background sets the background attribute.
func (e *BodyElement[C]) Background(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "background", Value: v})
	return e
}



// Onafterprint sets the onafterprint attribute.
func (e *BodyElement[C]) Onafterprint(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onafterprint", Value: v})
	return e
}



// Onbeforeprint sets the onbeforeprint attribute.
func (e *BodyElement[C]) Onbeforeprint(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onbeforeprint", Value: v})
	return e
}



// Onbeforeunload sets the onbeforeunload attribute.
func (e *BodyElement[C]) Onbeforeunload(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onbeforeunload", Value: v})
	return e
}



// Onhashchange sets the onhashchange attribute.
func (e *BodyElement[C]) Onhashchange(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onhashchange", Value: v})
	return e
}



// Onlanguagechange sets the onlanguagechange attribute.
func (e *BodyElement[C]) Onlanguagechange(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onlanguagechange", Value: v})
	return e
}



// Onmessage sets the onmessage attribute.
func (e *BodyElement[C]) Onmessage(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onmessage", Value: v})
	return e
}



// Onmessageerror sets the onmessageerror attribute.
func (e *BodyElement[C]) Onmessageerror(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onmessageerror", Value: v})
	return e
}



// Onoffline sets the onoffline attribute.
func (e *BodyElement[C]) Onoffline(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onoffline", Value: v})
	return e
}



// Ononline sets the ononline attribute.
func (e *BodyElement[C]) Ononline(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "ononline", Value: v})
	return e
}



// Onpagehide sets the onpagehide attribute.
func (e *BodyElement[C]) Onpagehide(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpagehide", Value: v})
	return e
}



// Onpagereveal sets the onpagereveal attribute.
func (e *BodyElement[C]) Onpagereveal(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpagereveal", Value: v})
	return e
}



// Onpageshow sets the onpageshow attribute.
func (e *BodyElement[C]) Onpageshow(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpageshow", Value: v})
	return e
}



// Onpageswap sets the onpageswap attribute.
func (e *BodyElement[C]) Onpageswap(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpageswap", Value: v})
	return e
}



// Onpopstate sets the onpopstate attribute.
func (e *BodyElement[C]) Onpopstate(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onpopstate", Value: v})
	return e
}



// Onrejectionhandled sets the onrejectionhandled attribute.
func (e *BodyElement[C]) Onrejectionhandled(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onrejectionhandled", Value: v})
	return e
}



// Onstorage sets the onstorage attribute.
func (e *BodyElement[C]) Onstorage(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onstorage", Value: v})
	return e
}



// Onunhandledrejection sets the onunhandledrejection attribute.
func (e *BodyElement[C]) Onunhandledrejection(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onunhandledrejection", Value: v})
	return e
}



// Onunload sets the onunload attribute.
func (e *BodyElement[C]) Onunload(v string) *BodyElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "onunload", Value: v})
	return e
}


