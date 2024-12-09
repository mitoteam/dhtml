package dhtml

import (
	"github.com/mitoteam/mttools"
)

type InputFormItem struct {
	FormItemExtBase
	inputType string
}

// force interfaces implementation
var _ FormItemExtI = (*InputFormItem)(nil)

func NewFormInput(name, inputType string) *InputFormItem {
	fi := &InputFormItem{inputType: inputType}
	fi.name = SafeId(name)
	fi.renderF = fi.Render

	return fi
}

func (fi *InputFormItem) Label(v any) *InputFormItem {
	fi.label.Append(v)
	return fi
}

func (fi *InputFormItem) DefaultValue(v any) *InputFormItem {
	fi.defaultValue = v
	return fi
}

func (fi *InputFormItem) Note(v any) *InputFormItem {
	fi.note.Append(v)
	return fi
}

func (fi *InputFormItem) Render() HtmlPiece {
	var out HtmlPiece

	if !fi.label.IsEmpty() {
		out.Append(
			NewTag("label").Attribute("for", fi.GetId()).Class("form-label").Append(fi.label),
		)
	}

	input_tag := NewTag("input").Id(fi.GetId()).Class("form-control").
		Attribute("type", fi.inputType).
		Attribute("value", mttools.AnyToString(fi.defaultValue))

	out.Append(input_tag)

	if !fi.note.IsEmpty() {
		out.Append(Div().Class("form-text").Append(fi.note))
	}

	return out
}
