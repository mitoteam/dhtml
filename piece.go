package dhtml

import (
	"fmt"
	"strings"
)

// HtmlPiece is set of one or several html elements (or no elements at all). Could be tags, complex elements, text content etc.
// Every HtmlPiece as an element itself (so it can be rendered as HTML).
type HtmlPiece struct {
	list []ElementI
}

// force interfaces implementation declaring fake variable
var _ ElementI = (*HtmlPiece)(nil)
var _ fmt.Stringer = (*HtmlPiece)(nil)

// Shorthand helper for NewHtmlPiece() constructor
func Piece(first_element any) *HtmlPiece {
	return NewHtmlPiece().AppendElement(AnyToElement(first_element))
}

// Actual Constructor
func NewHtmlPiece() *HtmlPiece {
	return &HtmlPiece{
		list: make([]ElementI, 0),
	}
}

func (p *HtmlPiece) IsEmpty() bool {
	return len(p.list) == 0
}

// Adds something to list
func (p *HtmlPiece) Append(v any) *HtmlPiece {
	if v == nil || v == "" {
		//nothing to append
		return p
	}

	switch v := v.(type) {
	case HtmlPiece:
		return p.AppendPiece(&v)

	case *HtmlPiece:
		return p.AppendPiece(v)

	case ElementI:
		return p.AppendElement(v)

	default:
		return p.AppendElement(AnyToElement(v))
	}
}

// Adds single element to list
func (p *HtmlPiece) AppendElement(e ElementI) *HtmlPiece {
	p.list = append(p.list, e)

	return p
}

// Adds single element to list
func (p *HtmlPiece) AppendPiece(another_piece *HtmlPiece) *HtmlPiece {
	p.list = append(p.list, another_piece.list...)

	return p
}

// Adds text element to list
func (p *HtmlPiece) AppendText(text string) *HtmlPiece {
	p.list = append(p.list, Text(text))

	return p
}

func (p *HtmlPiece) Textf(format string, a ...any) *HtmlPiece {
	p.list = append(p.list, Textf(format, a...))

	return p
}

// Elements count
func (p *HtmlPiece) GetElementsCount() int {
	return len(p.list)
}

// ElementI implementation
func (p *HtmlPiece) GetTags() TagsList {
	tag_list := make(TagsList, 0)

	for _, element := range p.list {
		tag_list = append(tag_list, element.GetTags()...)
	}

	return tag_list
}

// render everything to string as HTML
func (p HtmlPiece) String() string {
	var sb strings.Builder

	for _, element := range p.list {
		for _, tag := range element.GetTags() {
			sb.WriteString(tag.String())
		}
	}

	return sb.String()
}

//=========== NamedHtmlPieces ================

// Set of named html pieces
type NamedHtmlPieces struct {
	pieces map[string]*HtmlPiece
}

func NewNamedHtmlPieces() NamedHtmlPieces {
	ps := NamedHtmlPieces{
		pieces: make(map[string]*HtmlPiece, 0),
	}

	return ps
}

func (np NamedHtmlPieces) Add(name string, v any) {
	if _, ok := np.pieces[name]; ok {
		np.pieces[name].Append(v)
	} else {
		np.pieces[name] = Piece(v)
	}
}

func (np NamedHtmlPieces) Set(name string, v any) {
	switch v := v.(type) {
	case HtmlPiece:
		np.pieces[name] = &v
	case *HtmlPiece:
		np.pieces[name] = v
	default:
		np.pieces[name] = Piece(v)
	}
}

func (np NamedHtmlPieces) GetOk(name string) (p *HtmlPiece, ok bool) {
	p, ok = np.pieces[name]
	return p, ok
}

func (np NamedHtmlPieces) Get(name string) *HtmlPiece {
	if p, ok := np.GetOk(name); ok {
		return p
	}

	return NewHtmlPiece() //empty piece
}
