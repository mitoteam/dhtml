package dhtml

import (
	"github.com/mitoteam/mttools"
)

type FormInputElement struct {
	FormItemExtBase
	tag        *Tag
	labelAfter bool // true = render <label> after <input>
}

// force interfaces implementation
var _ FormItemExtI = (*FormInputElement)(nil)

func NewFormInput(name, inputType string) *FormInputElement {
	e := &FormInputElement{
		tag: NewTag("input").Attribute("type", inputType),
	}

	e.name = SafeId(name)
	e.wrapped = true
	e.renderF = e.renderInput

	return e
}

func (fi *FormInputElement) Label(v any) *FormInputElement {
	fi.label.Append(v)
	return fi
}

func (fi *FormInputElement) Class(v any) *FormInputElement {
	fi.tag.Class(v)
	return fi
}

func (fi *FormInputElement) WrapperClass(v any) *FormInputElement {
	fi.FormItemBase.WrapperClass(v)
	return fi
}

func (fi *FormInputElement) DefaultValue(v any) *FormInputElement {
	fi.defaultValue = v
	return fi
}

func (fi *FormInputElement) Placeholder(s string) *FormInputElement {
	fi.tag.Attribute("placeholder", s)
	return fi
}

func (fi *FormInputElement) LabelAfter(b bool) *FormInputElement {
	fi.labelAfter = b
	return fi
}

func (fi *FormInputElement) Note(v any) *FormInputElement {
	fi.note.Append(v)
	return fi
}

func (e *FormInputElement) renderInput() (out HtmlPiece) {
	if !e.labelAfter {
		out.Append(e.renderLabel())
	}

	e.tag.Id(e.GetId()).Class("form-control").
		Attribute("name", e.GetName())

	if s := mttools.AnyToString(e.GetValue()); s != "" {
		e.tag.Attribute("value", s)
	}

	out.Append(e.tag)

	if e.labelAfter {
		out.Append(e.renderLabel())
	}

	if !e.note.IsEmpty() {
		out.Append(Div().Class("form-text").Append(e.note))
	}

	return out
}

func (e *FormInputElement) renderLabel() (out HtmlPiece) {
	if !e.label.IsEmpty() {
		out.Append(
			NewTag("label").Attribute("for", e.GetId()).Class("form-label").Append(e.label),
		)
	}

	return out
}

// ======================== checkbox =========================

func NewFormCheckbox(name string) *FormInputElement {
	e := NewFormInput(name, "checkbox")
	e.renderF = e.renderCheckbox
	return e
}

func (e *FormInputElement) renderCheckbox() (out HtmlPiece) {
	e.tag.Id(e.GetId()).
		Attribute("name", e.GetName()).
		Attribute("value", "on") // set value implicitly (see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox#value fo details)

	if !mttools.IsEmpty(e.GetValue()) {
		e.tag.Attribute("checked", "checked")
	}

	out.Append(e.tag)

	out.Append(e.renderLabel())

	if !e.note.IsEmpty() {
		out.Append(Div().Class("form-text").Append(e.note))
	}

	return out
}
