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

// Returns true if there is a region with `name`.
func (np *NamedHtmlPieces) Has(name string) bool {
	_, ok := np.pieces[name]
	return ok
}

// Adds `v` piece to `name` region. If there is no such region, it is created. If the region already exists, `v` is appended to it.
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

// Replaces `name` region with `v` piece. If there is no such region, it is created.
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

// Deletes `name` region. If there is no such region, nothing happens.
func (np *NamedHtmlPieces) Delete(name string) {
	delete(np.pieces, name)
}

func (np *NamedHtmlPieces) Clear() {
	maps.Clear(np.pieces)
}
