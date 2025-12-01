package html

import "fmt"

// PreElement represents the <pre> HTML element.

type PreElement[C HTMLContent] basicElement[C]


// Pre creates a new <pre> element.
func Pre(children ...HTMLContent) *PreElement[HTMLContent] {
	return &PreElement[HTMLContent]{
		tagName: "pre",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *PreElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *PreElement[C]) Attr(key, value string) *PreElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Width sets the width attribute.
func (e *PreElement[C]) Width(v int) *PreElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: fmt.Sprintf("%d", v)})
	return e
}


