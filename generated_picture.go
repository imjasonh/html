package html

// PictureElement represents the <picture> HTML element.

type PictureElement[C HTMLContent] basicElement[C]


// Picture creates a new <picture> element.
func Picture(children ...HTMLContent) *PictureElement[HTMLContent] {
	return &PictureElement[HTMLContent]{
		tagName: "picture",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *PictureElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *PictureElement[C]) Attr(key, value string) *PictureElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}

