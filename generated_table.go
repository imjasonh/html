package html

// TableElement represents the <table> HTML element.

type TableElement[C HTMLContent] basicElement[C]


// Table creates a new <table> element.
func Table(children ...HTMLContent) *TableElement[HTMLContent] {
	return &TableElement[HTMLContent]{
		tagName: "table",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *TableElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TableElement[C]) Attr(key, value string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *TableElement[C]) Align(v string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Border sets the border attribute.
func (e *TableElement[C]) Border(v string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "border", Value: v})
	return e
}



// Frame sets the frame attribute.
func (e *TableElement[C]) Frame(v string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "frame", Value: v})
	return e
}



// Rows sets the rows attribute.
func (e *TableElement[C]) Rows(v string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "rows", Value: v})
	return e
}



// Rules sets the rules attribute.
func (e *TableElement[C]) Rules(v string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "rules", Value: v})
	return e
}



// Summary sets the summary attribute.
func (e *TableElement[C]) Summary(v string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "summary", Value: v})
	return e
}



// TBodies sets the tBodies attribute.
func (e *TableElement[C]) TBodies(v string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "tBodies", Value: v})
	return e
}



// Width sets the width attribute.
func (e *TableElement[C]) Width(v string) *TableElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: v})
	return e
}


