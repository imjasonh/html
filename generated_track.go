package html

// TrackElement represents the <track> HTML element.

type TrackElement basicElement[HTMLContent]


// Track creates a new <track> element.
func Track() *TrackElement {
	return &TrackElement{
		tagName: "track",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *TrackElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *TrackElement) Attr(key, value string) *TrackElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Default sets the default attribute.
func (e *TrackElement) Default(v bool) *TrackElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "default", Value: "default"})
	}
	return e
}



// Kind sets the kind attribute.
func (e *TrackElement) Kind(v string) *TrackElement {
	e.attributes = append(e.attributes, Attribute{Key: "kind", Value: v})
	return e
}



// Label sets the label attribute.
func (e *TrackElement) Label(v string) *TrackElement {
	e.attributes = append(e.attributes, Attribute{Key: "label", Value: v})
	return e
}



// Src sets the src attribute.
func (e *TrackElement) Src(v string) *TrackElement {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}



// Srclang sets the srclang attribute.
func (e *TrackElement) Srclang(v string) *TrackElement {
	e.attributes = append(e.attributes, Attribute{Key: "srclang", Value: v})
	return e
}



// Track sets the track attribute.
func (e *TrackElement) Track(v string) *TrackElement {
	e.attributes = append(e.attributes, Attribute{Key: "track", Value: v})
	return e
}


