# html

Type-safe HTML generation for Go, generated from WebIDL.

## Usage

```go
import . "github.com/imjasonh/html"

page := Div(
    P(Raw("Hello, "), T(username), Raw("!")),
    A(Raw("Click here")).Href("/home"),
    Img().Src("/logo.png").Alt("Logo"),
).Render()
```

## Features

- Generated from [WHATWG HTML Standard](https://html.spec.whatwg.org/) WebIDL
- Type-safe attributes (bool, int, string)
- XSS protection: `T()` escapes user input, `Raw()` for trusted HTML
- `Attr(key, value)` escape hatch for arbitrary attributes
- Automatic HTML escaping of attribute values

## Generating

```bash
go run cmd/generate/main.go  # Fetches latest WebIDL and regenerates elements
```
