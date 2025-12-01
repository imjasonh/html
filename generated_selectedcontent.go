package html

// SelectedContentElement represents the <selectedcontent> HTML element.

type SelectedContentElement[C HTMLContent] basicElement[C]


// Selectedcontent creates a new <selectedcontent> element.
func Selectedcontent(children ...HTMLContent) *SelectedContentElement[HTMLContent] {
	return &SelectedContentElement[HTMLContent]{
		tagName: "selectedcontent",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *SelectedContentElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *SelectedContentElement[C]) Attr(key, value string) *SelectedContentElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}

