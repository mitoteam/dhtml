package dhtml

import (
	"runtime"
)

type (
	DebugElement struct {
		label HtmlPiece
		body  HtmlPiece
	}
)

// force interface implementation declaring fake variable
var _ ElementI = (*DebugElement)(nil)

func NewDebugElement(skip int) *DebugElement {
	e := &DebugElement{}

	if _, file, line, ok := runtime.Caller(skip); ok {
		e.label.Textf("%s:%d", file, line)
	}

	return e
}

func (e *DebugElement) Append(v any) *DebugElement {
	e.body.Append(v)
	return e
}

func (e *DebugElement) Textf(format string, a ...any) *DebugElement {
	e.body.Textf(format, a...)
	return e
}

func (e *DebugElement) GetTags() TagsList {
	div := Div().Attribute("style", "border: 1px solid red; padding: 10px; margin 10px;")

	if !e.label.IsEmpty() {
		div.Append(
			Div().Attribute("style", "font-weight: bold; text-align: right;").
				Append(NewTag("small").Append(e.label)),
		)
	}

	div.Append(e.body)

	return TagsList{div}
}
