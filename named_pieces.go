package dhtml

import (
	"github.com/mitoteam/mttools"
	"golang.org/x/exp/maps"
)

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

func (np *NamedHtmlPieces) Add(name string, v any) {
	if mttools.IsEmpty(v) {
		return //nothing to add
	}

	if _, ok := np.pieces[name]; ok {
		np.pieces[name].Append(v)
	} else {
		np.pieces[name] = Piece(v)
	}
}

func (np *NamedHtmlPieces) Set(name string, v any) {
	switch v := v.(type) {
	case HtmlPiece:
		np.pieces[name] = &v
	case *HtmlPiece:
		np.pieces[name] = v
	default:
		np.pieces[name] = Piece(v)
	}
}

func (np *NamedHtmlPieces) GetOk(name string) (p *HtmlPiece, ok bool) {
	p, ok = np.pieces[name]
	return p, ok
}

func (np *NamedHtmlPieces) Get(name string) *HtmlPiece {
	if p, ok := np.GetOk(name); ok {
		return p
	}

	return NewHtmlPiece() //empty piece
}

func (np *NamedHtmlPieces) IsEmpty(name string) bool {
	if p, ok := np.pieces[name]; ok {
		return p.IsEmpty()
	}

	return false
}

func (np *NamedHtmlPieces) Clear() {
	maps.Clear(np.pieces)
}
