package dhtml

import (
	"fmt"
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

func Comment(text string) *Tag {
	return &Tag{
		kind: tagKindComment,
		text: text,
	}
}

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

func RenderValueE(title, value HtmlPiece, emptyLabel string) *Tag {
	valueP := Piece(value)

	if valueP.IsEmpty() {
		valueP.Append("[" + emptyLabel + "]")
	}

	return RenderValue(title, value)
}

func Dbg(format string, a ...any) ElementI {
	e := NewDebugElement(2)

	e.Textf(format, a...)

	return e
}
