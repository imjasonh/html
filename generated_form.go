package html

// FormElement represents the <form> HTML element.

type FormElement[C HTMLContent] basicElement[C]


// Form creates a new <form> element.
func Form(children ...HTMLContent) *FormElement[HTMLContent] {
	return &FormElement[HTMLContent]{
		tagName: "form",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *FormElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *FormElement[C]) Attr(key, value string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// AcceptCharset sets the acceptCharset attribute.
func (e *FormElement[C]) AcceptCharset(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "acceptCharset", Value: v})
	return e
}



// Action sets the action attribute.
func (e *FormElement[C]) Action(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "action", Value: v})
	return e
}



// Autocomplete sets the autocomplete attribute.
func (e *FormElement[C]) Autocomplete(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "autocomplete", Value: v})
	return e
}



// Elements sets the elements attribute.
func (e *FormElement[C]) Elements(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "elements", Value: v})
	return e
}



// Encoding sets the encoding attribute.
func (e *FormElement[C]) Encoding(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "encoding", Value: v})
	return e
}



// Enctype sets the enctype attribute.
func (e *FormElement[C]) Enctype(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "enctype", Value: v})
	return e
}



// Method sets the method attribute.
func (e *FormElement[C]) Method(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "method", Value: v})
	return e
}



// Name sets the name attribute.
func (e *FormElement[C]) Name(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// NoValidate sets the noValidate attribute.
func (e *FormElement[C]) NoValidate(v bool) *FormElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "noValidate", Value: "noValidate"})
	}
	return e
}



// Rel sets the rel attribute.
func (e *FormElement[C]) Rel(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "rel", Value: v})
	return e
}



// RelList sets the relList attribute.
func (e *FormElement[C]) RelList(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "relList", Value: v})
	return e
}



// Target sets the target attribute.
func (e *FormElement[C]) Target(v string) *FormElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "target", Value: v})
	return e
}


