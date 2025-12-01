package html

// CanvasElement represents the <canvas> HTML element.

type CanvasElement[C HTMLContent] basicElement[C]


// Canvas creates a new <canvas> element.
func Canvas(children ...HTMLContent) *CanvasElement[HTMLContent] {
	return &CanvasElement[HTMLContent]{
		tagName: "canvas",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *CanvasElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *CanvasElement[C]) Attr(key, value string) *CanvasElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}

