package dhtml

import (
	"fmt"

	"github.com/mitoteam/mttools"
)

//Some basic type and helper shorthands.

// Function returning some html.
type RenderFunc func() HtmlPiece

func Div() *Tag {
	return NewTag("div")
}

func Span() *Tag {
	return NewTag("span")
}

func Text(text string) *Tag {
	return &Tag{
		kind: tagKindText,
		text: text,
	}
}

func Textf(format string, args ...any) *Tag {
	return &Tag{
		kind: tagKindText,
		text: fmt.Sprintf(format, args...),
	}
}

func UnsafeText(text string) *Tag {
	return &Tag{
		kind: tagKindUnsafeText,
		text: text,
	}
}

func Comment(text string) *Tag {
	return &Tag{
		kind: tagKindComment,
		text: text,
	}
}

func EmptyLabel(label string) *Tag {
	span := Span()
	settings.EmptyLabelRendererF(label, span)
	return span
}

// Renders title and some value
func RenderValue(title, value any) *Tag {
	tag := Div()

	titleP := Piece(title)
	valueP := Piece(value)

	if !titleP.IsEmpty() {
		tag.Append(NewTag("strong").Append(titleP)).Append(": ")
	}

	tag.Append(valueP)

	return tag
}

// Renders title and some value. If value is empty, render EmptyLabel instead
func RenderValueE(title, value any, emptyLabel string) *Tag {
	var valueP *HtmlPiece

	if mttools.IsEmpty(value) {
		valueP = Piece(EmptyLabel(emptyLabel))
	} else {
		valueP = Piece(value)

		if valueP.IsEmpty() {
			valueP.Append(EmptyLabel(emptyLabel))
		}
	}

	return RenderValue(title, valueP)
}

func Dbg(format string, a ...any) ElementI {
	e := NewDebugElement(2)

	e.Textf(format, a...)

	return e
}
