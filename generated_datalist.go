package html

// DataListElement represents the <datalist> HTML element.

type DataListElement[C HTMLContent] basicElement[C]


// Datalist creates a new <datalist> element.
func Datalist(children ...HTMLContent) *DataListElement[HTMLContent] {
	return &DataListElement[HTMLContent]{
		tagName: "datalist",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *DataListElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *DataListElement[C]) Attr(key, value string) *DataListElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Options sets the options attribute.
func (e *DataListElement[C]) Options(v string) *DataListElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "options", Value: v})
	return e
}


