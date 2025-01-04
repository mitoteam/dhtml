package dhtml

import (
	"fmt"
	"strings"
)

// HtmlPiece is set of one or several html elements (or no elements at all). Could be tags, complex elements, text content etc.
// Every HtmlPiece as an element itself (so it can be rendered as HTML).
type HtmlPiece struct {
	list    []ElementI // added elements, tags or another pieces
	tagList TagList    // cached rendered contents
}

// force interfaces implementation declaring fake variable
var _ ElementI = (*HtmlPiece)(nil)
var _ fmt.Stringer = (*HtmlPiece)(nil)

// If firstElement is HtmlPiece, return it.
// Else create new HtmlPiece and add firstElement to its contents.
func Piece(firstElement any) *HtmlPiece {
	switch v := firstElement.(type) {
	case HtmlPiece:
		return &v

	case *HtmlPiece:
		return v

	default:
		return NewHtmlPiece().AppendElement(AnyToElement(firstElement))
	}
}

// Actual Constructor
func NewHtmlPiece() *HtmlPiece {
	return &HtmlPiece{
		list: make([]ElementI, 0),
	}
}

// Adds something to the piece: another piece, ElemenetI, any string, Stringer or other value.
func (p *HtmlPiece) Append(v ...any) *HtmlPiece {
	for _, v := range v {
		if v == nil || v == "" {
			//nothing to append
			continue
		}

		switch v := v.(type) {
		case HtmlPiece:
			p.AppendPiece(&v)

		case *HtmlPiece:
			p.AppendPiece(v)

		case ElementI:
			p.AppendElement(v)

		default:
			p.AppendElement(AnyToElement(v))
		}
	}

	return p
}

// Adds single element
func (p *HtmlPiece) AppendElement(e ElementI) *HtmlPiece {
	p.list = append(p.list, e)

	return p
}

// Adds another piece elements to this one
func (p *HtmlPiece) AppendPiece(another_piece *HtmlPiece) *HtmlPiece {
	if another_piece.IsEmpty() {
		//nothing to add
		return p
	}

	p.list = append(p.list, another_piece.list...)

	return p
}

// Adds text element to piece
func (p *HtmlPiece) AppendText(text string) *HtmlPiece {
	p.list = append(p.list, Text(text))

	return p
}

// Format and add text element
func (p *HtmlPiece) Textf(format string, a ...any) *HtmlPiece {
	p.list = append(p.list, Textf(format, a...))

	return p
}

// Returns true if piece has no anything added to it
func (p *HtmlPiece) IsEmpty() bool {
	return p.Len() == 0
}

// Elements count
func (p *HtmlPiece) Len() int {
	return len(p.list)
}

// remove all cached contents
func (p *HtmlPiece) Clear() *HtmlPiece {
	p.tagList = make(TagList, 0)
	return p
}

// Calls f function for each element.
func (p *HtmlPiece) Walk(f ElementWalkFunc, args ...any) {
	if len(p.tagList) > 0 {
		// already rendered
		for _, e := range p.tagList {
			f(e, args...)
		}
	} else {
		for _, e := range p.list {
			f(e, args...)
		}
	}
}

// Calls f function for each element with recursion.
func (p *HtmlPiece) WalkR(f ElementWalkFunc, args ...any) {
	p.Walk(func(e ElementI, args ...any) {
		f(e, args...)

		//and dive deeper
		switch v := e.(type) {
		case *HtmlPiece:
			v.WalkR(f, args...)
		case *Tag:
			v.WalkR(f, args...)
		}
	})
}

// ElementI implementation
func (p *HtmlPiece) GetTags() TagList {
	//p.tagList - cached tags if it was already rendered
	if len(p.tagList) == 0 {
		for _, element := range p.list {
			p.tagList = append(p.tagList, element.GetTags()...)
		}
	}

	return p.tagList
}

// Render everything to string as HTML
func (p *HtmlPiece) String() string {
	var sb strings.Builder

	for _, tag := range p.GetTags() {
		sb.WriteString(tag.String())
	}

	return sb.String()
}

// Render just text content only
func (p *HtmlPiece) RawString() string {
	var sb strings.Builder

	for _, tag := range p.GetTags() {
		if tag.kind == tagKindText {
			sb.WriteString(tag.text)
		}
	}

	return sb.String()
}
