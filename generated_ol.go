package html

import "fmt"

// OListElement represents the <ol> HTML element.

type OListElement[C HTMLContent] basicElement[C]


// Ol creates a new <ol> element.
func Ol(children ...HTMLContent) *OListElement[HTMLContent] {
	return &OListElement[HTMLContent]{
		tagName: "ol",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *OListElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *OListElement[C]) Attr(key, value string) *OListElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Compact sets the compact attribute.
func (e *OListElement[C]) Compact(v bool) *OListElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "compact", Value: "compact"})
	}
	return e
}



// Reversed sets the reversed attribute.
func (e *OListElement[C]) Reversed(v bool) *OListElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "reversed", Value: "reversed"})
	}
	return e
}



// Start sets the start attribute.
func (e *OListElement[C]) Start(v int) *OListElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "start", Value: fmt.Sprintf("%d", v)})
	return e
}



// Type sets the type attribute.
func (e *OListElement[C]) Type(v string) *OListElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}


