package html

// AreaElement represents the <area> HTML element.

type AreaElement basicElement[HTMLContent]


// Area creates a new <area> element.
func Area() *AreaElement {
	return &AreaElement{
		tagName: "area",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *AreaElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *AreaElement) Attr(key, value string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Alt sets the alt attribute.
func (e *AreaElement) Alt(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "alt", Value: v})
	return e
}



// Coords sets the coords attribute.
func (e *AreaElement) Coords(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "coords", Value: v})
	return e
}



// Download sets the download attribute.
func (e *AreaElement) Download(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "download", Value: v})
	return e
}



// Hash sets the hash attribute.
func (e *AreaElement) Hash(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "hash", Value: v})
	return e
}



// Host sets the host attribute.
func (e *AreaElement) Host(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "host", Value: v})
	return e
}



// Hostname sets the hostname attribute.
func (e *AreaElement) Hostname(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "hostname", Value: v})
	return e
}



// Href sets the href attribute.
func (e *AreaElement) Href(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "href", Value: v})
	return e
}



// NoHref sets the noHref attribute.
func (e *AreaElement) NoHref(v bool) *AreaElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "noHref", Value: "noHref"})
	}
	return e
}



// Origin sets the origin attribute.
func (e *AreaElement) Origin(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "origin", Value: v})
	return e
}



// Password sets the password attribute.
func (e *AreaElement) Password(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "password", Value: v})
	return e
}



// Pathname sets the pathname attribute.
func (e *AreaElement) Pathname(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "pathname", Value: v})
	return e
}



// Ping sets the ping attribute.
func (e *AreaElement) Ping(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "ping", Value: v})
	return e
}



// Port sets the port attribute.
func (e *AreaElement) Port(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "port", Value: v})
	return e
}



// Protocol sets the protocol attribute.
func (e *AreaElement) Protocol(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "protocol", Value: v})
	return e
}



// ReferrerPolicy sets the referrerPolicy attribute.
func (e *AreaElement) ReferrerPolicy(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "referrerPolicy", Value: v})
	return e
}



// Rel sets the rel attribute.
func (e *AreaElement) Rel(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "rel", Value: v})
	return e
}



// RelList sets the relList attribute.
func (e *AreaElement) RelList(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "relList", Value: v})
	return e
}



// Search sets the search attribute.
func (e *AreaElement) Search(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "search", Value: v})
	return e
}



// Shape sets the shape attribute.
func (e *AreaElement) Shape(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "shape", Value: v})
	return e
}



// Target sets the target attribute.
func (e *AreaElement) Target(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "target", Value: v})
	return e
}



// Username sets the username attribute.
func (e *AreaElement) Username(v string) *AreaElement {
	e.attributes = append(e.attributes, Attribute{Key: "username", Value: v})
	return e
}


