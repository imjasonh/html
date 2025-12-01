package html

// StyleElement represents the <style> HTML element.

type StyleElement[C HTMLContent] basicElement[C]


// Style creates a new <style> element.
func Style(children ...HTMLContent) *StyleElement[HTMLContent] {
	return &StyleElement[HTMLContent]{
		tagName: "style",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *StyleElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *StyleElement[C]) Attr(key, value string) *StyleElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Blocking sets the blocking attribute.
func (e *StyleElement[C]) Blocking(v string) *StyleElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "blocking", Value: v})
	return e
}



// Disabled sets the disabled attribute.
func (e *StyleElement[C]) Disabled(v bool) *StyleElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// Media sets the media attribute.
func (e *StyleElement[C]) Media(v string) *StyleElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "media", Value: v})
	return e
}



// Type sets the type attribute.
func (e *StyleElement[C]) Type(v string) *StyleElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}


