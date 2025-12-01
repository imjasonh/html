package html

import "fmt"

// MeterElement represents the <meter> HTML element.

type MeterElement[C HTMLContent] basicElement[C]


// Meter creates a new <meter> element.
func Meter(children ...HTMLContent) *MeterElement[HTMLContent] {
	return &MeterElement[HTMLContent]{
		tagName: "meter",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *MeterElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *MeterElement[C]) Attr(key, value string) *MeterElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// High sets the high attribute.
func (e *MeterElement[C]) High(v float64) *MeterElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "high", Value: fmt.Sprintf("%g", v)})
	return e
}



// Labels sets the labels attribute.
func (e *MeterElement[C]) Labels(v string) *MeterElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "labels", Value: v})
	return e
}



// Low sets the low attribute.
func (e *MeterElement[C]) Low(v float64) *MeterElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "low", Value: fmt.Sprintf("%g", v)})
	return e
}



// Max sets the max attribute.
func (e *MeterElement[C]) Max(v float64) *MeterElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "max", Value: fmt.Sprintf("%g", v)})
	return e
}



// Min sets the min attribute.
func (e *MeterElement[C]) Min(v float64) *MeterElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "min", Value: fmt.Sprintf("%g", v)})
	return e
}



// Optimum sets the optimum attribute.
func (e *MeterElement[C]) Optimum(v float64) *MeterElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "optimum", Value: fmt.Sprintf("%g", v)})
	return e
}



// Value sets the value attribute.
func (e *MeterElement[C]) Value(v float64) *MeterElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "value", Value: fmt.Sprintf("%g", v)})
	return e
}


