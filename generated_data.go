package html

// DataElement represents the <data> HTML element.

type DataElement[C HTMLContent] basicElement[C]


// Data creates a new <data> element.
func Data(children ...HTMLContent) *DataElement[HTMLContent] {
	return &DataElement[HTMLContent]{
		tagName: "data",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *DataElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *DataElement[C]) Attr(key, value string) *DataElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Value sets the value attribute.
func (e *DataElement[C]) Value(v string) *DataElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: v})
	return e
}


