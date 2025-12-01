package html

import "fmt"

// ProgressElement represents the <progress> HTML element.

type ProgressElement[C HTMLContent] basicElement[C]


// Progress creates a new <progress> element.
func Progress(children ...HTMLContent) *ProgressElement[HTMLContent] {
	return &ProgressElement[HTMLContent]{
		tagName: "progress",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *ProgressElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *ProgressElement[C]) Attr(key, value string) *ProgressElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Labels sets the labels attribute.
func (e *ProgressElement[C]) Labels(v string) *ProgressElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "labels", Value: v})
	return e
}



// Max sets the max attribute.
func (e *ProgressElement[C]) Max(v float64) *ProgressElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "max", Value: fmt.Sprintf("%g", v)})
	return e
}



// Position sets the position attribute.
func (e *ProgressElement[C]) Position(v float64) *ProgressElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "position", Value: fmt.Sprintf("%g", v)})
	return e
}



// Value sets the value attribute.
func (e *ProgressElement[C]) Value(v float64) *ProgressElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: fmt.Sprintf("%g", v)})
	return e
}


