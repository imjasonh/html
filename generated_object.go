package html

// ObjectElement represents the <object> HTML element.

type ObjectElement[C HTMLContent] basicElement[C]


// Object creates a new <object> element.
func Object(children ...HTMLContent) *ObjectElement[HTMLContent] {
	return &ObjectElement[HTMLContent]{
		tagName: "object",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *ObjectElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *ObjectElement[C]) Attr(key, value string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Align sets the align attribute.
func (e *ObjectElement[C]) Align(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Archive sets the archive attribute.
func (e *ObjectElement[C]) Archive(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "archive", Value: v})
	return e
}



// Code sets the code attribute.
func (e *ObjectElement[C]) Code(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "code", Value: v})
	return e
}



// CodeBase sets the codeBase attribute.
func (e *ObjectElement[C]) CodeBase(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "codeBase", Value: v})
	return e
}



// CodeType sets the codeType attribute.
func (e *ObjectElement[C]) CodeType(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "codeType", Value: v})
	return e
}



// Data sets the data attribute.
func (e *ObjectElement[C]) Data(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "data", Value: v})
	return e
}



// Declare sets the declare attribute.
func (e *ObjectElement[C]) Declare(v bool) *ObjectElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "declare", Value: "declare"})
	}
	return e
}



// Height sets the height attribute.
func (e *ObjectElement[C]) Height(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "height", Value: v})
	return e
}



// Name sets the name attribute.
func (e *ObjectElement[C]) Name(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Standby sets the standby attribute.
func (e *ObjectElement[C]) Standby(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "standby", Value: v})
	return e
}



// Type sets the type attribute.
func (e *ObjectElement[C]) Type(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// UseMap sets the useMap attribute.
func (e *ObjectElement[C]) UseMap(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "useMap", Value: v})
	return e
}



// ValidationMessage sets the validationMessage attribute.
func (e *ObjectElement[C]) ValidationMessage(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validationMessage", Value: v})
	return e
}



// Validity sets the validity attribute.
func (e *ObjectElement[C]) Validity(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validity", Value: v})
	return e
}



// Width sets the width attribute.
func (e *ObjectElement[C]) Width(v string) *ObjectElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: v})
	return e
}



// WillValidate sets the willValidate attribute.
func (e *ObjectElement[C]) WillValidate(v bool) *ObjectElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "willValidate", Value: "willValidate"})
	}
	return e
}


