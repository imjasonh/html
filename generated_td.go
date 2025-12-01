package html

import "fmt"

// TableCellElement represents the <td> HTML element.

type TableCellElement[C HTMLContent] basicElement[C]


// Td creates a new <td> element.
func Td(children ...HTMLContent) *TableCellElement[HTMLContent] {
	return &TableCellElement[HTMLContent]{
		tagName: "td",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *TableCellElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TableCellElement[C]) Attr(key, value string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Abbr sets the abbr attribute.
func (e *TableCellElement[C]) Abbr(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "abbr", Value: v})
	return e
}



// Align sets the align attribute.
func (e *TableCellElement[C]) Align(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Axis sets the axis attribute.
func (e *TableCellElement[C]) Axis(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "axis", Value: v})
	return e
}



// CellIndex sets the cellIndex attribute.
func (e *TableCellElement[C]) CellIndex(v int) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "cellIndex", Value: fmt.Sprintf("%d", v)})
	return e
}



// Ch sets the ch attribute.
func (e *TableCellElement[C]) Ch(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "ch", Value: v})
	return e
}



// ChOff sets the chOff attribute.
func (e *TableCellElement[C]) ChOff(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "chOff", Value: v})
	return e
}



// Headers sets the headers attribute.
func (e *TableCellElement[C]) Headers(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "headers", Value: v})
	return e
}



// Height sets the height attribute.
func (e *TableCellElement[C]) Height(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "height", Value: v})
	return e
}



// NoWrap sets the noWrap attribute.
func (e *TableCellElement[C]) NoWrap(v bool) *TableCellElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "noWrap", Value: "noWrap"})
	}
	return e
}



// Scope sets the scope attribute.
func (e *TableCellElement[C]) Scope(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "scope", Value: v})
	return e
}



// VAlign sets the vAlign attribute.
func (e *TableCellElement[C]) VAlign(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "vAlign", Value: v})
	return e
}



// Width sets the width attribute.
func (e *TableCellElement[C]) Width(v string) *TableCellElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: v})
	return e
}


