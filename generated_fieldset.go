package html

// FieldSetElement represents the <fieldset> HTML element.

type FieldSetElement[C HTMLContent] basicElement[C]


// Fieldset creates a new <fieldset> element.
func Fieldset(children ...HTMLContent) *FieldSetElement[HTMLContent] {
	return &FieldSetElement[HTMLContent]{
		tagName: "fieldset",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *FieldSetElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *FieldSetElement[C]) Attr(key, value string) *FieldSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Disabled sets the disabled attribute.
func (e *FieldSetElement[C]) Disabled(v bool) *FieldSetElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// Elements sets the elements attribute.
func (e *FieldSetElement[C]) Elements(v string) *FieldSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "elements", Value: v})
	return e
}



// Name sets the name attribute.
func (e *FieldSetElement[C]) Name(v string) *FieldSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Type sets the type attribute.
func (e *FieldSetElement[C]) Type(v string) *FieldSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// ValidationMessage sets the validationMessage attribute.
func (e *FieldSetElement[C]) ValidationMessage(v string) *FieldSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validationMessage", Value: v})
	return e
}



// Validity sets the validity attribute.
func (e *FieldSetElement[C]) Validity(v string) *FieldSetElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validity", Value: v})
	return e
}



// WillValidate sets the willValidate attribute.
func (e *FieldSetElement[C]) WillValidate(v bool) *FieldSetElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "willValidate", Value: "willValidate"})
	}
	return e
}


