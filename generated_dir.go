package html

// DirectoryElement represents the <dir> HTML element.

type DirectoryElement[C HTMLContent] basicElement[C]


// Dir creates a new <dir> element.
func Dir(children ...HTMLContent) *DirectoryElement[HTMLContent] {
	return &DirectoryElement[HTMLContent]{
		tagName: "dir",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *DirectoryElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *DirectoryElement[C]) Attr(key, value string) *DirectoryElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Compact sets the compact attribute.
func (e *DirectoryElement[C]) Compact(v bool) *DirectoryElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "compact", Value: "compact"})
	}
	return e
}


