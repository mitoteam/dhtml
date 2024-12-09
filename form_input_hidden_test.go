package dhtml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHidden(t *testing.T) {
	e := NewFormHidden("hname", "hvalue")
	html := "<div class=\"form-item\">\n  <input name=\"hname\" type=\"hidden\" value=\"hvalue\" /></div>"

	require.Equal(t, html, Piece(e).String())
}
