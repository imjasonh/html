package html

// TemplateElement represents the <template> HTML element.

type TemplateElement[C HTMLContent] basicElement[C]


// Template creates a new <template> element.
func Template(children ...HTMLContent) *TemplateElement[HTMLContent] {
	return &TemplateElement[HTMLContent]{
		tagName: "template",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *TemplateElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TemplateElement[C]) Attr(key, value string) *TemplateElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Content sets the content attribute.
func (e *TemplateElement[C]) Content(v string) *TemplateElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "content", Value: v})
	return e
}



// ShadowRootClonable sets the shadowRootClonable attribute.
func (e *TemplateElement[C]) ShadowRootClonable(v bool) *TemplateElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "shadowRootClonable", Value: "shadowRootClonable"})
	}
	return e
}



// ShadowRootCustomElementRegistry sets the shadowRootCustomElementRegistry attribute.
func (e *TemplateElement[C]) ShadowRootCustomElementRegistry(v string) *TemplateElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "shadowRootCustomElementRegistry", Value: v})
	return e
}



// ShadowRootDelegatesFocus sets the shadowRootDelegatesFocus attribute.
func (e *TemplateElement[C]) ShadowRootDelegatesFocus(v bool) *TemplateElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "shadowRootDelegatesFocus", Value: "shadowRootDelegatesFocus"})
	}
	return e
}



// ShadowRootMode sets the shadowRootMode attribute.
func (e *TemplateElement[C]) ShadowRootMode(v string) *TemplateElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "shadowRootMode", Value: v})
	return e
}



// ShadowRootSerializable sets the shadowRootSerializable attribute.
func (e *TemplateElement[C]) ShadowRootSerializable(v bool) *TemplateElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "shadowRootSerializable", Value: "shadowRootSerializable"})
	}
	return e
}


