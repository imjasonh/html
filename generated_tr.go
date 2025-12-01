package html

import "fmt"

// TableRowElement represents the <tr> HTML element.

type TableRowElement[C HTMLContent] basicElement[C]


// Tr creates a new <tr> element.
func Tr(children ...HTMLContent) *TableRowElement[HTMLContent] {
	return &TableRowElement[HTMLContent]{
		tagName: "tr",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *TableRowElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TableRowElement[C]) Attr(key, value string) *TableRowElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *TableRowElement[C]) Align(v string) *TableRowElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Cells sets the cells attribute.
func (e *TableRowElement[C]) Cells(v string) *TableRowElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "cells", Value: v})
	return e
}



// Ch sets the ch attribute.
func (e *TableRowElement[C]) Ch(v string) *TableRowElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "ch", Value: v})
	return e
}



// ChOff sets the chOff attribute.
func (e *TableRowElement[C]) ChOff(v string) *TableRowElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "chOff", Value: v})
	return e
}



// RowIndex sets the rowIndex attribute.
func (e *TableRowElement[C]) RowIndex(v int) *TableRowElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "rowIndex", Value: fmt.Sprintf("%d", v)})
	return e
}



// SectionRowIndex sets the sectionRowIndex attribute.
func (e *TableRowElement[C]) SectionRowIndex(v int) *TableRowElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "sectionRowIndex", Value: fmt.Sprintf("%d", v)})
	return e
}



// VAlign sets the vAlign attribute.
func (e *TableRowElement[C]) VAlign(v string) *TableRowElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "vAlign", Value: v})
	return e
}


