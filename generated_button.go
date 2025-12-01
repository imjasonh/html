package html

// ButtonElement represents the <button> HTML element.

type ButtonElement[C HTMLContent] basicElement[C]


// Button creates a new <button> element.
func Button(children ...HTMLContent) *ButtonElement[HTMLContent] {
	return &ButtonElement[HTMLContent]{
		tagName: "button",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *ButtonElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *ButtonElement[C]) Attr(key, value string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Command sets the command attribute.
func (e *ButtonElement[C]) Command(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "command", Value: v})
	return e
}



// Disabled sets the disabled attribute.
func (e *ButtonElement[C]) Disabled(v bool) *ButtonElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// FormAction sets the formAction attribute.
func (e *ButtonElement[C]) FormAction(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "formAction", Value: v})
	return e
}



// FormEnctype sets the formEnctype attribute.
func (e *ButtonElement[C]) FormEnctype(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "formEnctype", Value: v})
	return e
}



// FormMethod sets the formMethod attribute.
func (e *ButtonElement[C]) FormMethod(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "formMethod", Value: v})
	return e
}



// FormNoValidate sets the formNoValidate attribute.
func (e *ButtonElement[C]) FormNoValidate(v bool) *ButtonElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "formNoValidate", Value: "formNoValidate"})
	}
	return e
}



// FormTarget sets the formTarget attribute.
func (e *ButtonElement[C]) FormTarget(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "formTarget", Value: v})
	return e
}



// Labels sets the labels attribute.
func (e *ButtonElement[C]) Labels(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "labels", Value: v})
	return e
}



// Name sets the name attribute.
func (e *ButtonElement[C]) Name(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// PopoverTargetAction sets the popoverTargetAction attribute.
func (e *ButtonElement[C]) PopoverTargetAction(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "popoverTargetAction", Value: v})
	return e
}



// Type sets the type attribute.
func (e *ButtonElement[C]) Type(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// ValidationMessage sets the validationMessage attribute.
func (e *ButtonElement[C]) ValidationMessage(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validationMessage", Value: v})
	return e
}



// Validity sets the validity attribute.
func (e *ButtonElement[C]) Validity(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "validity", Value: v})
	return e
}



// Value sets the value attribute.
func (e *ButtonElement[C]) Value(v string) *ButtonElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: v})
	return e
}



// WillValidate sets the willValidate attribute.
func (e *ButtonElement[C]) WillValidate(v bool) *ButtonElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "willValidate", Value: "willValidate"})
	}
	return e
}


