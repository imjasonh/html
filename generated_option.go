package html

import "fmt"

// OptionElement represents the <option> HTML element.

type OptionElement[C HTMLContent] basicElement[C]


// Option creates a new <option> element.
func Option(children ...HTMLContent) *OptionElement[HTMLContent] {
	return &OptionElement[HTMLContent]{
		tagName: "option",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *OptionElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *OptionElement[C]) Attr(key, value string) *OptionElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// DefaultSelected sets the defaultSelected attribute.
func (e *OptionElement[C]) DefaultSelected(v bool) *OptionElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "defaultSelected", Value: "defaultSelected"})
	}
	return e
}



// Disabled sets the disabled attribute.
func (e *OptionElement[C]) Disabled(v bool) *OptionElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// Index sets the index attribute.
func (e *OptionElement[C]) Index(v int) *OptionElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "index", Value: fmt.Sprintf("%d", v)})
	return e
}



// Label sets the label attribute.
func (e *OptionElement[C]) Label(v string) *OptionElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "label", Value: v})
	return e
}



// Selected sets the selected attribute.
func (e *OptionElement[C]) Selected(v bool) *OptionElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "selected", Value: "selected"})
	}
	return e
}



// Text sets the text attribute.
func (e *OptionElement[C]) Text(v string) *OptionElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "text", Value: v})
	return e
}



// Value sets the value attribute.
func (e *OptionElement[C]) Value(v string) *OptionElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: v})
	return e
}


