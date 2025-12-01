package html

import "fmt"

// InputElement represents the <input> HTML element.

type InputElement basicElement[HTMLContent]


// Input creates a new <input> element.
func Input() *InputElement {
	return &InputElement{
		tagName: "input",
		isVoid:  true,
		
	}
}

// Render returns the HTML string representation.
func (e *InputElement) Render() string {
	return (*basicElement[HTMLContent])(e).Render()
}

// Attr sets an arbitrary attribute as an escape hatch.
func (e *InputElement) Attr(key, value string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: key, Value: value})
	return e
}


// Accept sets the accept attribute.
func (e *InputElement) Accept(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "accept", Value: v})
	return e
}



// Align sets the align attribute.
func (e *InputElement) Align(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "align", Value: v})
	return e
}



// Alpha sets the alpha attribute.
func (e *InputElement) Alpha(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "alpha", Value: "alpha"})
	}
	return e
}



// Alt sets the alt attribute.
func (e *InputElement) Alt(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "alt", Value: v})
	return e
}



// Autocomplete sets the autocomplete attribute.
func (e *InputElement) Autocomplete(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "autocomplete", Value: v})
	return e
}



// Checked sets the checked attribute.
func (e *InputElement) Checked(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "checked", Value: "checked"})
	}
	return e
}



// ColorSpace sets the colorSpace attribute.
func (e *InputElement) ColorSpace(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "colorSpace", Value: v})
	return e
}



// DefaultChecked sets the defaultChecked attribute.
func (e *InputElement) DefaultChecked(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "defaultChecked", Value: "defaultChecked"})
	}
	return e
}



// DefaultValue sets the defaultValue attribute.
func (e *InputElement) DefaultValue(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "defaultValue", Value: v})
	return e
}



// DirName sets the dirName attribute.
func (e *InputElement) DirName(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "dirName", Value: v})
	return e
}



// Disabled sets the disabled attribute.
func (e *InputElement) Disabled(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "disabled", Value: "disabled"})
	}
	return e
}



// FormAction sets the formAction attribute.
func (e *InputElement) FormAction(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "formAction", Value: v})
	return e
}



// FormEnctype sets the formEnctype attribute.
func (e *InputElement) FormEnctype(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "formEnctype", Value: v})
	return e
}



// FormMethod sets the formMethod attribute.
func (e *InputElement) FormMethod(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "formMethod", Value: v})
	return e
}



// FormNoValidate sets the formNoValidate attribute.
func (e *InputElement) FormNoValidate(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "formNoValidate", Value: "formNoValidate"})
	}
	return e
}



// FormTarget sets the formTarget attribute.
func (e *InputElement) FormTarget(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "formTarget", Value: v})
	return e
}



// Indeterminate sets the indeterminate attribute.
func (e *InputElement) Indeterminate(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "indeterminate", Value: "indeterminate"})
	}
	return e
}



// Max sets the max attribute.
func (e *InputElement) Max(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "max", Value: v})
	return e
}



// MaxLength sets the maxLength attribute.
func (e *InputElement) MaxLength(v int) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "maxLength", Value: fmt.Sprintf("%d", v)})
	return e
}



// Min sets the min attribute.
func (e *InputElement) Min(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "min", Value: v})
	return e
}



// MinLength sets the minLength attribute.
func (e *InputElement) MinLength(v int) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "minLength", Value: fmt.Sprintf("%d", v)})
	return e
}



// Multiple sets the multiple attribute.
func (e *InputElement) Multiple(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "multiple", Value: "multiple"})
	}
	return e
}



// Name sets the name attribute.
func (e *InputElement) Name(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "name", Value: v})
	return e
}



// Pattern sets the pattern attribute.
func (e *InputElement) Pattern(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "pattern", Value: v})
	return e
}



// Placeholder sets the placeholder attribute.
func (e *InputElement) Placeholder(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "placeholder", Value: v})
	return e
}



// PopoverTargetAction sets the popoverTargetAction attribute.
func (e *InputElement) PopoverTargetAction(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "popoverTargetAction", Value: v})
	return e
}



// ReadOnly sets the readOnly attribute.
func (e *InputElement) ReadOnly(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "readOnly", Value: "readOnly"})
	}
	return e
}



// Required sets the required attribute.
func (e *InputElement) Required(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "required", Value: "required"})
	}
	return e
}



// Src sets the src attribute.
func (e *InputElement) Src(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "src", Value: v})
	return e
}



// Step sets the step attribute.
func (e *InputElement) Step(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "step", Value: v})
	return e
}



// Type sets the type attribute.
func (e *InputElement) Type(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "type", Value: v})
	return e
}



// UseMap sets the useMap attribute.
func (e *InputElement) UseMap(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "useMap", Value: v})
	return e
}



// ValidationMessage sets the validationMessage attribute.
func (e *InputElement) ValidationMessage(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "validationMessage", Value: v})
	return e
}



// Validity sets the validity attribute.
func (e *InputElement) Validity(v string) *InputElement {
	e.attributes = append(e.attributes, Attribute{Key: "validity", Value: v})
	return e
}



// WillValidate sets the willValidate attribute.
func (e *InputElement) WillValidate(v bool) *InputElement {
	if v {
		e.attributes = append(e.attributes, Attribute{Key: "willValidate", Value: "willValidate"})
	}
	return e
}


