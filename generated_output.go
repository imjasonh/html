package html

// OutputElement represents the <output> HTML element.

type OutputElement[C HTMLContent] basicElement[C]


// Output creates a new <output> element.
func Output(children ...HTMLContent) *OutputElement[HTMLContent] {
	return &OutputElement[HTMLContent]{
		tagName: "output",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *OutputElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *OutputElement[C]) Attr(key, value string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// DefaultValue sets the defaultValue attribute.
func (e *OutputElement[C]) DefaultValue(v string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "defaultValue", Value: v})
	return e
}



// HtmlFor sets the htmlFor attribute.
func (e *OutputElement[C]) HtmlFor(v string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "htmlFor", Value: v})
	return e
}



// Labels sets the labels attribute.
func (e *OutputElement[C]) Labels(v string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "labels", Value: v})
	return e
}



// Name sets the name attribute.
func (e *OutputElement[C]) Name(v string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Type sets the type attribute.
func (e *OutputElement[C]) Type(v string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// ValidationMessage sets the validationMessage attribute.
func (e *OutputElement[C]) ValidationMessage(v string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validationMessage", Value: v})
	return e
}



// Validity sets the validity attribute.
func (e *OutputElement[C]) Validity(v string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validity", Value: v})
	return e
}



// Value sets the value attribute.
func (e *OutputElement[C]) Value(v string) *OutputElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: v})
	return e
}



// WillValidate sets the willValidate attribute.
func (e *OutputElement[C]) WillValidate(v bool) *OutputElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "willValidate", Value: "willValidate"})
	}
	return e
}


