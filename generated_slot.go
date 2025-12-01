package html

// SlotElement represents the <slot> HTML element.

type SlotElement[C HTMLContent] basicElement[C]


// Slot creates a new <slot> element.
func Slot(children ...HTMLContent) *SlotElement[HTMLContent] {
	return &SlotElement[HTMLContent]{
		tagName: "slot",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *SlotElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *SlotElement[C]) Attr(key, value string) *SlotElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Name sets the name attribute.
func (e *SlotElement[C]) Name(v string) *SlotElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}


