package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

const webidlURL = "https://raw.githubusercontent.com/w3c/webref/main/ed/idl/html.idl"

// VoidElements lists HTML elements that are self-closing
var voidElements = map[string]bool{
	"area": true, "base": true, "br": true, "col": true, "embed": true,
	"hr": true, "img": true, "input": true, "link": true, "meta": true,
	"param": true, "source": true, "track": true, "wbr": true,
}

// AttributeType represents the Go type for an attribute
type AttributeType string

const (
	TypeString AttributeType = "string"
	TypeBool   AttributeType = "bool"
	TypeInt    AttributeType = "int"
	TypeFloat  AttributeType = "float64"
	TypeURL    AttributeType = "string" // URLs are strings
	TypeEnum   AttributeType = "string" // Enums are strings
)

// Attribute represents an HTML attribute with its type
type Attribute struct {
	Name       string
	GoName     string // Capitalized method name
	Type       AttributeType
	EnumValues []string // For enum types
}

// ElementDef defines an HTML element
type ElementDef struct {
	Name       string
	GoTypeName string // e.g., "DivElement"
	GoFuncName string // e.g., "Div"
	IsVoid     bool
	Attributes []Attribute
}

func main() {
	fmt.Println("Fetching WebIDL from", webidlURL)
	resp, err := http.Get(webidlURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching WebIDL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error: HTTP %d\n", resp.StatusCode)
		os.Exit(1)
	}

	idlContent, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading WebIDL: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Parsing WebIDL...")
	elements := parseWebIDL(string(idlContent))
	fmt.Printf("Found %d HTML elements\n", len(elements))

	fmt.Println("Generating Go files...")
	if err := generateElements(elements); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %d element files\n", len(elements))
}

func parseWebIDL(idl string) []ElementDef {
	elements := make(map[string]*ElementDef)
	mixins := make(map[string][]Attribute) // Store mixin attributes
	includes := make(map[string][]string)  // Track which interfaces include which mixins

	// Regex patterns for parsing
	interfacePattern := regexp.MustCompile(`(?m)^interface\s+(HTML\w+Element)\s*:\s*HTMLElement\s*\{`)
	partialPattern := regexp.MustCompile(`(?m)^partial\s+interface\s+(HTML\w+Element)\s*\{`)
	mixinPattern := regexp.MustCompile(`(?m)^interface\s+mixin\s+(\w+)\s*\{`)
	includesPattern := regexp.MustCompile(`(?m)^(HTML\w+Element)\s+includes\s+(\w+);`)
	attributePattern := regexp.MustCompile(`(?m)^\s*(?:\[[^\]]+\]\s+)*(?:stringifier\s+)?(?:readonly\s+)?attribute\s+(\w+)\s+(\w+);`)

	scanner := bufio.NewScanner(strings.NewReader(idl))
	var currentInterface string
	var currentMixin string
	var inInterface bool
	var inMixin bool
	braceDepth := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Track brace depth to know when we exit an interface
		braceDepth += strings.Count(line, "{") - strings.Count(line, "}")

		// Check for includes statements
		if matches := includesPattern.FindStringSubmatch(line); matches != nil {
			interfaceName := matches[1]
			mixinName := matches[2]
			includes[interfaceName] = append(includes[interfaceName], mixinName)
			continue
		}

		// Check for mixin definitions
		if matches := mixinPattern.FindStringSubmatch(line); matches != nil {
			currentMixin = matches[1]
			inMixin = true
			if mixins[currentMixin] == nil {
				mixins[currentMixin] = []Attribute{}
			}
			continue
		}

		// Check for interface or partial interface declarations
		if matches := interfacePattern.FindStringSubmatch(line); matches != nil {
			currentInterface = matches[1]
			inInterface = true

			// Extract element name (e.g., "HTMLDivElement" -> "div")
			elemName := extractElementName(currentInterface)
			if elemName != "" {
				if _, exists := elements[elemName]; !exists {
					elements[elemName] = &ElementDef{
						Name:       elemName,
						GoTypeName: strings.TrimPrefix(currentInterface, "HTML"),
						GoFuncName: capitalize(elemName),
						IsVoid:     voidElements[elemName],
						Attributes: []Attribute{},
					}
				}
			}
			continue
		}

		if matches := partialPattern.FindStringSubmatch(line); matches != nil {
			currentInterface = matches[1]
			inInterface = true
			continue
		}

		// Exit interface/mixin when braces close
		if braceDepth == 0 {
			inInterface = false
			inMixin = false
			currentInterface = ""
			currentMixin = ""
		}

		// Parse attributes within interface
		if inInterface && currentInterface != "" {
			if matches := attributePattern.FindStringSubmatch(line); matches != nil {
				attrType := matches[1]
				attrName := matches[2]

				elemName := extractElementName(currentInterface)
				if elemName != "" && elements[elemName] != nil {
					goType := mapIDLTypeToGo(attrType)
					attr := Attribute{
						Name:   attrName,
						GoName: capitalize(attrName),
						Type:   goType,
					}

					// Check if attribute already exists
					exists := false
					for _, existing := range elements[elemName].Attributes {
						if existing.Name == attrName {
							exists = true
							break
						}
					}
					if !exists {
						elements[elemName].Attributes = append(elements[elemName].Attributes, attr)
					}
				}
			}
		}

		// Parse attributes within mixin
		if inMixin && currentMixin != "" {
			if matches := attributePattern.FindStringSubmatch(line); matches != nil {
				attrType := matches[1]
				attrName := matches[2]

				goType := mapIDLTypeToGo(attrType)
				attr := Attribute{
					Name:   attrName,
					GoName: capitalize(attrName),
					Type:   goType,
				}

				// Check if attribute already exists
				exists := false
				for _, existing := range mixins[currentMixin] {
					if existing.Name == attrName {
						exists = true
						break
					}
				}
				if !exists {
					mixins[currentMixin] = append(mixins[currentMixin], attr)
				}
			}
		}
	}

	// Apply mixins to interfaces that include them
	for interfaceName, mixinNames := range includes {
		elemName := extractElementName(interfaceName)
		if elemName != "" && elements[elemName] != nil {
			for _, mixinName := range mixinNames {
				if mixinAttrs, ok := mixins[mixinName]; ok {
					for _, attr := range mixinAttrs {
						// Check if attribute already exists
						exists := false
						for _, existing := range elements[elemName].Attributes {
							if existing.Name == attr.Name {
								exists = true
								break
							}
						}
						if !exists {
							elements[elemName].Attributes = append(elements[elemName].Attributes, attr)
						}
					}
				}
			}
		}
	}

	// Convert map to sorted slice
	var result []ElementDef
	for _, elem := range elements {
		// Sort attributes
		sort.Slice(elem.Attributes, func(i, j int) bool {
			return elem.Attributes[i].Name < elem.Attributes[j].Name
		})
		result = append(result, *elem)
	}

	// Sort elements by name
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
}

func extractElementName(interfaceName string) string {
	// Extract element name from interface name
	// e.g., "HTMLDivElement" -> "div"
	// e.g., "HTMLAnchorElement" -> "a"

	if !strings.HasPrefix(interfaceName, "HTML") || !strings.HasSuffix(interfaceName, "Element") {
		return ""
	}

	name := strings.TrimPrefix(interfaceName, "HTML")
	name = strings.TrimSuffix(name, "Element")

	// Handle special cases
	specialCases := map[string]string{
		"Anchor":       "a",
		"DList":        "dl",
		"Directory":    "dir",
		"Heading":      "", // Skip generic heading
		"Image":        "img",
		"Mod":          "", // Skip generic mod (ins/del)
		"OList":        "ol",
		"Paragraph":    "p",
		"Quote":        "q",
		"TableCaption": "caption",
		"TableCell":    "td",
		"TableCol":     "col",
		"TableRow":     "tr",
		"TableSection": "", // Skip generic section
		"UList":        "ul",
	}

	if special, ok := specialCases[name]; ok {
		return special
	}

	return strings.ToLower(name)
}

func mapIDLTypeToGo(idlType string) AttributeType {
	// Map WebIDL types to Go types
	switch idlType {
	case "boolean":
		return TypeBool
	case "long", "unsigned long", "short", "unsigned short":
		return TypeInt
	case "double", "float":
		return TypeFloat
	case "DOMString", "USVString":
		return TypeString
	default:
		return TypeString // Default to string for unknown types
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	// Convert hyphenated names to camelCase (e.g., "http-equiv" -> "HttpEquiv")
	parts := strings.Split(s, "-")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

const elementTemplate = `package html
{{if needsFmt .Attributes}}
import "fmt"
{{end}}
// {{.GoTypeName}} represents the <{{.Name}}> HTML element.
{{if .IsVoid}}
type {{.GoTypeName}} basicElement[HTMLContent]
{{else}}
type {{.GoTypeName}}[C HTMLContent] basicElement[C]
{{end}}

// {{.GoFuncName}} creates a new <{{.Name}}> element.
func {{.GoFuncName}}({{if not .IsVoid}}children ...HTMLContent{{end}}) *{{.GoTypeName}}{{if not .IsVoid}}[HTMLContent]{{end}} {
	return &{{.GoTypeName}}{{if not .IsVoid}}[HTMLContent]{{end}}{
		tagName: "{{.Name}}",
		{{if .IsVoid}}isVoid:  true,{{end}}
		{{if not .IsVoid}}content: children,{{end}}
	}
}

// Render returns the HTML string representation.
func (e *{{.GoTypeName}}{{if not .IsVoid}}[C]{{end}}) Render() string {
	{{if .IsVoid}}return (*basicElement[HTMLContent])(e).Render(){{else}}return (*basicElement[C])(e).Render(){{end}}
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *{{.GoTypeName}}{{if not .IsVoid}}[C]{{end}}) Attr(key, value string) *{{.GoTypeName}}{{if not .IsVoid}}[C]{{end}} {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}
{{range .Attributes}}
{{if eq .Type "bool"}}
// {{.GoName}} sets the {{.Name}} attribute.
func (e *{{$.GoTypeName}}{{if not $.IsVoid}}[C]{{end}}) {{.GoName}}(v bool) *{{$.GoTypeName}}{{if not $.IsVoid}}[C]{{end}} {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "{{.Name}}", Value: "{{.Name}}"})
	}
	return e
}
{{else if eq .Type "int"}}
// {{.GoName}} sets the {{.Name}} attribute.
func (e *{{$.GoTypeName}}{{if not $.IsVoid}}[C]{{end}}) {{.GoName}}(v int) *{{$.GoTypeName}}{{if not $.IsVoid}}[C]{{end}} {
	e.attributes = append(e.attributes, Attribute{Key: "{{.Name}}", Value: fmt.Sprintf("%d", v)})
	return e
}
{{else if eq .Type "float64"}}
// {{.GoName}} sets the {{.Name}} attribute.
func (e *{{$.GoTypeName}}{{if not $.IsVoid}}[C]{{end}}) {{.GoName}}(v float64) *{{$.GoTypeName}}{{if not $.IsVoid}}[C]{{end}} {
	e.attributes = append(e.attributes, Attribute{Key: "{{.Name}}", Value: fmt.Sprintf("%g", v)})
	return e
}
{{else}}
// {{.GoName}} sets the {{.Name}} attribute.
func (e *{{$.GoTypeName}}{{if not $.IsVoid}}[C]{{end}}) {{.GoName}}(v string) *{{$.GoTypeName}}{{if not $.IsVoid}}[C]{{end}} {
	e.attributes = append(e.attributes, Attribute{Key: "{{.Name}}", Value: v})
	return e
}
{{end}}
{{end}}
`

func needsFmt(attrs []Attribute) bool {
	for _, attr := range attrs {
		if attr.Type == TypeInt || attr.Type == TypeFloat {
			return true
		}
	}
	return false
}

func generateElements(elements []ElementDef) error {
	tmpl := template.Must(template.New("element").Funcs(template.FuncMap{
		"needsFmt": needsFmt,
	}).Parse(elementTemplate))

	for _, elem := range elements {
		filename := fmt.Sprintf("generated_%s.go", elem.Name)
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("creating file %s: %w", filename, err)
		}

		if err := tmpl.Execute(f, elem); err != nil {
			f.Close()
			return fmt.Errorf("executing template for %s: %w", elem.Name, err)
		}
		f.Close()
	}

	return nil
}
