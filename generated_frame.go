package html

// FrameElement represents the <frame> HTML element.

type FrameElement[C HTMLContent] basicElement[C]


// Frame creates a new <frame> element.
func Frame(children ...HTMLContent) *FrameElement[HTMLContent] {
	return &FrameElement[HTMLContent]{
		tagName: "frame",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *FrameElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *FrameElement[C]) Attr(key, value string) *FrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// FrameBorder sets the frameBorder attribute.
func (e *FrameElement[C]) FrameBorder(v string) *FrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "frameBorder", Value: v})
	return e
}



// LongDesc sets the longDesc attribute.
func (e *FrameElement[C]) LongDesc(v string) *FrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "longDesc", Value: v})
	return e
}



// Name sets the name attribute.
func (e *FrameElement[C]) Name(v string) *FrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// NoResize sets the noResize attribute.
func (e *FrameElement[C]) NoResize(v bool) *FrameElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "noResize", Value: "noResize"})
	}
	return e
}



// Scrolling sets the scrolling attribute.
func (e *FrameElement[C]) Scrolling(v string) *FrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "scrolling", Value: v})
	return e
}



// Src sets the src attribute.
func (e *FrameElement[C]) Src(v string) *FrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}


