package html

import "fmt"

// LIElement represents the <li> HTML element.

type LIElement[C HTMLContent] basicElement[C]


// Li creates a new <li> element.
func Li(children ...HTMLContent) *LIElement[HTMLContent] {
	return &LIElement[HTMLContent]{
		tagName: "li",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *LIElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *LIElement[C]) Attr(key, value string) *LIElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Type sets the type attribute.
func (e *LIElement[C]) Type(v string) *LIElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// Value sets the value attribute.
func (e *LIElement[C]) Value(v int) *LIElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: fmt.Sprintf("%d", v)})
	return e
}


