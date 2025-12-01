package html

import "fmt"

// TextAreaElement represents the <textarea> HTML element.

type TextAreaElement[C HTMLContent] basicElement[C]


// Textarea creates a new <textarea> element.
func Textarea(children ...HTMLContent) *TextAreaElement[HTMLContent] {
	return &TextAreaElement[HTMLContent]{
		tagName: "textarea",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *TextAreaElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TextAreaElement[C]) Attr(key, value string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Autocomplete sets the autocomplete attribute.
func (e *TextAreaElement[C]) Autocomplete(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "autocomplete", Value: v})
	return e
}



// DefaultValue sets the defaultValue attribute.
func (e *TextAreaElement[C]) DefaultValue(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "defaultValue", Value: v})
	return e
}



// DirName sets the dirName attribute.
func (e *TextAreaElement[C]) DirName(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "dirName", Value: v})
	return e
}



// Disabled sets the disabled attribute.
func (e *TextAreaElement[C]) Disabled(v bool) *TextAreaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// Labels sets the labels attribute.
func (e *TextAreaElement[C]) Labels(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "labels", Value: v})
	return e
}



// MaxLength sets the maxLength attribute.
func (e *TextAreaElement[C]) MaxLength(v int) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "maxLength", Value: fmt.Sprintf("%d", v)})
	return e
}



// MinLength sets the minLength attribute.
func (e *TextAreaElement[C]) MinLength(v int) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "minLength", Value: fmt.Sprintf("%d", v)})
	return e
}



// Name sets the name attribute.
func (e *TextAreaElement[C]) Name(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Placeholder sets the placeholder attribute.
func (e *TextAreaElement[C]) Placeholder(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "placeholder", Value: v})
	return e
}



// ReadOnly sets the readOnly attribute.
func (e *TextAreaElement[C]) ReadOnly(v bool) *TextAreaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "readOnly", Value: "readOnly"})
	}
	return e
}



// Required sets the required attribute.
func (e *TextAreaElement[C]) Required(v bool) *TextAreaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "required", Value: "required"})
	}
	return e
}



// SelectionDirection sets the selectionDirection attribute.
func (e *TextAreaElement[C]) SelectionDirection(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "selectionDirection", Value: v})
	return e
}



// Type sets the type attribute.
func (e *TextAreaElement[C]) Type(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// ValidationMessage sets the validationMessage attribute.
func (e *TextAreaElement[C]) ValidationMessage(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validationMessage", Value: v})
	return e
}



// Validity sets the validity attribute.
func (e *TextAreaElement[C]) Validity(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validity", Value: v})
	return e
}



// WillValidate sets the willValidate attribute.
func (e *TextAreaElement[C]) WillValidate(v bool) *TextAreaElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "willValidate", Value: "willValidate"})
	}
	return e
}



// Wrap sets the wrap attribute.
func (e *TextAreaElement[C]) Wrap(v string) *TextAreaElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "wrap", Value: v})
	return e
}


