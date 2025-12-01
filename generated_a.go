package html

// AnchorElement represents the <a> HTML element.

type AnchorElement[C HTMLContent] basicElement[C]


// A creates a new <a> element.
func A(children ...HTMLContent) *AnchorElement[HTMLContent] {
	return &AnchorElement[HTMLContent]{
		tagName: "a",
		
		content: children,
	}
}

// Render returns the HTML string representation.
func (e *AnchorElement[C]) Render() string {
	return (*basicElement[C])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *AnchorElement[C]) Attr(key, value string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Charset sets the charset attribute.
func (e *AnchorElement[C]) Charset(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "charset", Value: v})
	return e
}



// Coords sets the coords attribute.
func (e *AnchorElement[C]) Coords(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "coords", Value: v})
	return e
}



// Download sets the download attribute.
func (e *AnchorElement[C]) Download(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "download", Value: v})
	return e
}



// Hash sets the hash attribute.
func (e *AnchorElement[C]) Hash(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "hash", Value: v})
	return e
}



// Host sets the host attribute.
func (e *AnchorElement[C]) Host(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "host", Value: v})
	return e
}



// Hostname sets the hostname attribute.
func (e *AnchorElement[C]) Hostname(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "hostname", Value: v})
	return e
}



// Href sets the href attribute.
func (e *AnchorElement[C]) Href(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "href", Value: v})
	return e
}



// Hreflang sets the hreflang attribute.
func (e *AnchorElement[C]) Hreflang(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "hreflang", Value: v})
	return e
}



// Name sets the name attribute.
func (e *AnchorElement[C]) Name(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Origin sets the origin attribute.
func (e *AnchorElement[C]) Origin(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "origin", Value: v})
	return e
}



// Password sets the password attribute.
func (e *AnchorElement[C]) Password(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "password", Value: v})
	return e
}



// Pathname sets the pathname attribute.
func (e *AnchorElement[C]) Pathname(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "pathname", Value: v})
	return e
}



// Ping sets the ping attribute.
func (e *AnchorElement[C]) Ping(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "ping", Value: v})
	return e
}



// Port sets the port attribute.
func (e *AnchorElement[C]) Port(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "port", Value: v})
	return e
}



// Protocol sets the protocol attribute.
func (e *AnchorElement[C]) Protocol(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "protocol", Value: v})
	return e
}



// ReferrerPolicy sets the referrerPolicy attribute.
func (e *AnchorElement[C]) ReferrerPolicy(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "referrerPolicy", Value: v})
	return e
}



// Rel sets the rel attribute.
func (e *AnchorElement[C]) Rel(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "rel", Value: v})
	return e
}



// RelList sets the relList attribute.
func (e *AnchorElement[C]) RelList(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "relList", Value: v})
	return e
}



// Rev sets the rev attribute.
func (e *AnchorElement[C]) Rev(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "rev", Value: v})
	return e
}



// Search sets the search attribute.
func (e *AnchorElement[C]) Search(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "search", Value: v})
	return e
}



// Shape sets the shape attribute.
func (e *AnchorElement[C]) Shape(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "shape", Value: v})
	return e
}



// Target sets the target attribute.
func (e *AnchorElement[C]) Target(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "target", Value: v})
	return e
}



// Text sets the text attribute.
func (e *AnchorElement[C]) Text(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "text", Value: v})
	return e
}



// Type sets the type attribute.
func (e *AnchorElement[C]) Type(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// Username sets the username attribute.
func (e *AnchorElement[C]) Username(v string) *AnchorElement[C] {
	e.attributes = append(e.attributes, Attribute{Key: "username", Value: v})
	return e
}


