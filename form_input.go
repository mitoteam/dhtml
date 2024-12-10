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
	fi := &FormInputElement{
		tag: NewTag("input").Attribute("type", inputType).Class("form-control"),
	}

	fi.name = SafeId(name)
	fi.wrapped = true
	fi.renderF = fi.Render

	return fi
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

func (fi *FormInputElement) Render() (out HtmlPiece) {
	if !fi.labelAfter {
		out.Append(fi.renderLabel())
	}

	fi.tag.Id(fi.GetId()).
		Attribute("name", fi.GetName())

	if s := mttools.AnyToString(fi.GetValue()); s != "" {
		fi.tag.Attribute("value", s)
	}

	out.Append(fi.tag)

	if fi.labelAfter {
		out.Append(fi.renderLabel())
	}

	if !fi.note.IsEmpty() {
		out.Append(Div().Class("form-text").Append(fi.note))
	}

	return out
}

func (fi *FormInputElement) renderLabel() (out HtmlPiece) {
	if !fi.label.IsEmpty() {
		out.Append(
			NewTag("label").Attribute("for", fi.GetId()).Class("form-label").Append(fi.label),
		)
	}

	return out
}
