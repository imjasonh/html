package html

// MenuElement represents the <menu> HTML element.

type MenuElement[C HTMLContent] basicElement[C]


// Menu creates a new <menu> element.
func Menu(children ...HTMLContent) *MenuElement[HTMLContent] {
	return &MenuElement[HTMLContent]{
		tagName: "menu",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *MenuElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *MenuElement[C]) Attr(key, value string) *MenuElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Compact sets the compact attribute.
func (e *MenuElement[C]) Compact(v bool) *MenuElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "compact", Value: "compact"})
	}
	return e
}


