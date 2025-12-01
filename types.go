package html

import "html"

// --- 1. Content Safety and Interfaces ---

// HTMLContent is the core interface for anything that can be rendered as part of the document.
type HTMLContent interface {
	// Render returns the HTML string representation of the content.
	Render() string
}

// Text represents escaped text content.
type Text struct {
	Value string
}

// Render automatically escapes the text content to prevent XSS.
func (t Text) Render() string {
	return html.EscapeString(t.Value)
}

// T is a helper function that escapes text for safe rendering.
func T(s string) Text {
	return Text{Value: s}
}

// Raw represents trusted HTML that should not be escaped.
// Use this ONLY for known-safe content from trusted sources.
type Raw string

// Render returns the raw HTML without escaping.
func (r Raw) Render() string {
	return string(r)
}

// --- 2. Element Structure and Generics ---

// Attribute represents a single key/value HTML attribute for internal storage.
type Attribute struct {
	Key   string
	Value string
}

// basicElement is the concrete base struct used by all WebIDL-mapped elements.
// The generic constraint C enforces the element's allowed content model.
type basicElement[C HTMLContent] struct {
	tagName    string
	attributes []Attribute
	content    []C
	isVoid     bool // True for self-closing tags like <img> or <br>
}

// Render iterates over attributes and content to build the final HTML string.
func (e *basicElement[C]) Render() string {
	var attrStr string
	// Concatenate attributes, escaping values.
	for _, attr := range e.attributes {
		attrStr += " " + attr.Key + "=\"" + html.EscapeString(attr.Value) + "\""
	}

	if e.isVoid {
		// Void elements are self-closing
		return "<" + e.tagName + attrStr + " />"
	}

	var contentStr string
	// Recursively render content.
	for _, content := range e.content {
		contentStr += content.Render()
	}

	return "<" + e.tagName + attrStr + ">" + contentStr + "</" + e.tagName + ">"
}
