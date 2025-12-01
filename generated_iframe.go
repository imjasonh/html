package html

// IFrameElement represents the <iframe> HTML element.

type IFrameElement[C HTMLContent] basicElement[C]


// Iframe creates a new <iframe> element.
func Iframe(children ...HTMLContent) *IFrameElement[HTMLContent] {
	return &IFrameElement[HTMLContent]{
		tagName: "iframe",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *IFrameElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *IFrameElement[C]) Attr(key, value string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *IFrameElement[C]) Align(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Allow sets the allow attribute.
func (e *IFrameElement[C]) Allow(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "allow", Value: v})
	return e
}



// AllowFullscreen sets the allowFullscreen attribute.
func (e *IFrameElement[C]) AllowFullscreen(v bool) *IFrameElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "allowFullscreen", Value: "allowFullscreen"})
	}
	return e
}



// FrameBorder sets the frameBorder attribute.
func (e *IFrameElement[C]) FrameBorder(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "frameBorder", Value: v})
	return e
}



// Height sets the height attribute.
func (e *IFrameElement[C]) Height(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "height", Value: v})
	return e
}



// Loading sets the loading attribute.
func (e *IFrameElement[C]) Loading(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "loading", Value: v})
	return e
}



// LongDesc sets the longDesc attribute.
func (e *IFrameElement[C]) LongDesc(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "longDesc", Value: v})
	return e
}



// Name sets the name attribute.
func (e *IFrameElement[C]) Name(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// ReferrerPolicy sets the referrerPolicy attribute.
func (e *IFrameElement[C]) ReferrerPolicy(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "referrerPolicy", Value: v})
	return e
}



// Sandbox sets the sandbox attribute.
func (e *IFrameElement[C]) Sandbox(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "sandbox", Value: v})
	return e
}



// Scrolling sets the scrolling attribute.
func (e *IFrameElement[C]) Scrolling(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "scrolling", Value: v})
	return e
}



// Src sets the src attribute.
func (e *IFrameElement[C]) Src(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}



// Width sets the width attribute.
func (e *IFrameElement[C]) Width(v string) *IFrameElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: v})
	return e
}


