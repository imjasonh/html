package html

import "fmt"

// SelectElement represents the <select> HTML element.

type SelectElement[C HTMLContent] basicElement[C]


// Select creates a new <select> element.
func Select(children ...HTMLContent) *SelectElement[HTMLContent] {
	return &SelectElement[HTMLContent]{
		tagName: "select",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *SelectElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *SelectElement[C]) Attr(key, value string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Autocomplete sets the autocomplete attribute.
func (e *SelectElement[C]) Autocomplete(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "autocomplete", Value: v})
	return e
}



// Disabled sets the disabled attribute.
func (e *SelectElement[C]) Disabled(v bool) *SelectElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// Labels sets the labels attribute.
func (e *SelectElement[C]) Labels(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "labels", Value: v})
	return e
}



// Multiple sets the multiple attribute.
func (e *SelectElement[C]) Multiple(v bool) *SelectElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "multiple", Value: "multiple"})
	}
	return e
}



// Name sets the name attribute.
func (e *SelectElement[C]) Name(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Options sets the options attribute.
func (e *SelectElement[C]) Options(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "options", Value: v})
	return e
}



// Required sets the required attribute.
func (e *SelectElement[C]) Required(v bool) *SelectElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "required", Value: "required"})
	}
	return e
}



// SelectedIndex sets the selectedIndex attribute.
func (e *SelectElement[C]) SelectedIndex(v int) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "selectedIndex", Value: fmt.Sprintf("%d", v)})
	return e
}



// SelectedOptions sets the selectedOptions attribute.
func (e *SelectElement[C]) SelectedOptions(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "selectedOptions", Value: v})
	return e
}



// Type sets the type attribute.
func (e *SelectElement[C]) Type(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// ValidationMessage sets the validationMessage attribute.
func (e *SelectElement[C]) ValidationMessage(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validationMessage", Value: v})
	return e
}



// Validity sets the validity attribute.
func (e *SelectElement[C]) Validity(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validity", Value: v})
	return e
}



// Value sets the value attribute.
func (e *SelectElement[C]) Value(v string) *SelectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: v})
	return e
}



// WillValidate sets the willValidate attribute.
func (e *SelectElement[C]) WillValidate(v bool) *SelectElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "willValidate", Value: "willValidate"})
	}
	return e
}


