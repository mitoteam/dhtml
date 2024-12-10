package dhtml

import (
	"github.com/mitoteam/mttools"
)

type InputFormItem struct {
	FormItemExtBase
	tag        *Tag
	labelAfter bool // true = render <label> after <input>
}

// force interfaces implementation
var _ FormItemExtI = (*InputFormItem)(nil)

func NewFormInput(name, inputType string) *InputFormItem {
	fi := &InputFormItem{
		tag: NewTag("input").Attribute("type", inputType).Class("form-control"),
	}

	fi.name = SafeId(name)
	fi.wrapped = true
	fi.renderF = fi.Render

	return fi
}

func (fi *InputFormItem) Label(v any) *InputFormItem {
	fi.label.Append(v)
	return fi
}

func (fi *InputFormItem) Class(v any) *InputFormItem {
	fi.tag.Class(v)
	return fi
}

func (fi *InputFormItem) WrapperClass(v any) *InputFormItem {
	fi.FormItemBase.WrapperClass(v)
	return fi
}

func (fi *InputFormItem) DefaultValue(v any) *InputFormItem {
	fi.defaultValue = v
	return fi
}

func (fi *InputFormItem) Placeholder(s string) *InputFormItem {
	fi.tag.Attribute("placeholder", s)
	return fi
}

func (fi *InputFormItem) LabelAfter(b bool) *InputFormItem {
	fi.labelAfter = b
	return fi
}

func (fi *InputFormItem) Note(v any) *InputFormItem {
	fi.note.Append(v)
	return fi
}

func (fi *InputFormItem) Render() (out HtmlPiece) {
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

func (fi *InputFormItem) renderLabel() (out HtmlPiece) {
	if !fi.label.IsEmpty() {
		out.Append(
			NewTag("label").Attribute("for", fi.GetId()).Class("form-label").Append(fi.label),
		)
	}

	return out
}
