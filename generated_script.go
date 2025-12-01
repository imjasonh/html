package html

// ScriptElement represents the <script> HTML element.

type ScriptElement[C HTMLContent] basicElement[C]


// Script creates a new <script> element.
func Script(children ...HTMLContent) *ScriptElement[HTMLContent] {
	return &ScriptElement[HTMLContent]{
		tagName: "script",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *ScriptElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *ScriptElement[C]) Attr(key, value string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Async sets the async attribute.
func (e *ScriptElement[C]) Async(v bool) *ScriptElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "async", Value: "async"})
	}
	return e
}



// Blocking sets the blocking attribute.
func (e *ScriptElement[C]) Blocking(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "blocking", Value: v})
	return e
}



// Charset sets the charset attribute.
func (e *ScriptElement[C]) Charset(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "charset", Value: v})
	return e
}



// Defer sets the defer attribute.
func (e *ScriptElement[C]) Defer(v bool) *ScriptElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "defer", Value: "defer"})
	}
	return e
}



// Event sets the event attribute.
func (e *ScriptElement[C]) Event(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "event", Value: v})
	return e
}



// FetchPriority sets the fetchPriority attribute.
func (e *ScriptElement[C]) FetchPriority(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "fetchPriority", Value: v})
	return e
}



// HtmlFor sets the htmlFor attribute.
func (e *ScriptElement[C]) HtmlFor(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "htmlFor", Value: v})
	return e
}



// Integrity sets the integrity attribute.
func (e *ScriptElement[C]) Integrity(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "integrity", Value: v})
	return e
}



// NoModule sets the noModule attribute.
func (e *ScriptElement[C]) NoModule(v bool) *ScriptElement[C] {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "noModule", Value: "noModule"})
	}
	return e
}



// ReferrerPolicy sets the referrerPolicy attribute.
func (e *ScriptElement[C]) ReferrerPolicy(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "referrerPolicy", Value: v})
	return e
}



// Src sets the src attribute.
func (e *ScriptElement[C]) Src(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}



// Text sets the text attribute.
func (e *ScriptElement[C]) Text(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "text", Value: v})
	return e
}



// Type sets the type attribute.
func (e *ScriptElement[C]) Type(v string) *ScriptElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}


