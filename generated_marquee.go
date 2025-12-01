package html

import "fmt"

// MarqueeElement represents the <marquee> HTML element.

type MarqueeElement[C HTMLContent] basicElement[C]


// Marquee creates a new <marquee> element.
func Marquee(children ...HTMLContent) *MarqueeElement[HTMLContent] {
	return &MarqueeElement[HTMLContent]{
		tagName: "marquee",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *MarqueeElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *MarqueeElement[C]) Attr(key, value string) *MarqueeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Behavior sets the behavior attribute.
func (e *MarqueeElement[C]) Behavior(v string) *MarqueeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "behavior", Value: v})
	return e
}



// BgColor sets the bgColor attribute.
func (e *MarqueeElement[C]) BgColor(v string) *MarqueeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "bgColor", Value: v})
	return e
}



// Direction sets the direction attribute.
func (e *MarqueeElement[C]) Direction(v string) *MarqueeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "direction", Value: v})
	return e
}



// Height sets the height attribute.
func (e *MarqueeElement[C]) Height(v string) *MarqueeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "height", Value: v})
	return e
}



// Loop sets the loop attribute.
func (e *MarqueeElement[C]) Loop(v int) *MarqueeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "loop", Value: fmt.Sprintf("%d", v)})
	return e
}



// TrueSpeed sets the trueSpeed attribute.
func (e *MarqueeElement[C]) TrueSpeed(v bool) *MarqueeElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "trueSpeed", Value: "trueSpeed"})
	}
	return e
}



// Width sets the width attribute.
func (e *MarqueeElement[C]) Width(v string) *MarqueeElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "width", Value: v})
	return e
}


