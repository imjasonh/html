package html

// TimeElement represents the <time> HTML element.

type TimeElement[C HTMLContent] basicElement[C]


// Time creates a new <time> element.
func Time(children ...HTMLContent) *TimeElement[HTMLContent] {
	return &TimeElement[HTMLContent]{
		tagName: "time",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *TimeElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TimeElement[C]) Attr(key, value string) *TimeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// DateTime sets the dateTime attribute.
func (e *TimeElement[C]) DateTime(v string) *TimeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "dateTime", Value: v})
	return e
}


