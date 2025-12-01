package main

import (
	"fmt"

	. "github.com/imjasonh/html"
)

func main() {
	// Render to HTML string
	fmt.Println(Div(
		P(
			Raw("Welcome to the app. Here's a link: "),
			A(Raw("Click Here")).Href("/home").Target("_blank"),
			Raw("!"),
		).Attr("data-section", "intro"),

		// A self-closing element (void element)
		Img().
			Src("/asset/img.jpg").
			Alt("A Placeholder Image").
			Attr("loading", "lazy").
			Attr("data-category", "hero"),

		// XSS check: User input must be wrapped with T() for escaping
		P(T("User Input: <script>alert('pwned')</script>")),

		// More examples
		Ul(
			Li(Raw("Automatic XSS escaping")),
			Li(Raw("Type-safe attributes from WebIDL")),
			Li(Raw("Builder pattern with method chaining")),
		),

		Form(
			Label(Raw("Name:")),
			Input().
				Type("text").
				Name("username").
				Placeholder("Enter your name").
				Required(true),
			Button(Raw("Submit")).
				Type("submit"),
		).
			Action("/submit").
			Method("post"),

		Raw("<!-- This is a comment -->"), // Raw HTML insertion

		Raw(`<script src="i-want-this-to-be-raw.js"></script>`), // Raw HTML insertion
	).Render())
}
