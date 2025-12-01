package html

// ParamElement represents the <param> HTML element.

type ParamElement basicElement[HTMLContent]


// Param creates a new <param> element.
func Param() *ParamElement {
	return &ParamElement{
		tagName: "param",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *ParamElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *ParamElement) Attr(key, value string) *ParamElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Name sets the name attribute.
func (e *ParamElement) Name(v string) *ParamElement {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Type sets the type attribute.
func (e *ParamElement) Type(v string) *ParamElement {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// Value sets the value attribute.
func (e *ParamElement) Value(v string) *ParamElement {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: v})
	return e
}



// ValueType sets the valueType attribute.
func (e *ParamElement) ValueType(v string) *ParamElement {
	e.attributes = append(e.attributes, Attribute{Key: "valueType", Value: v})
	return e
}


