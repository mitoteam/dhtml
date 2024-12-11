package dhtml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHidden(t *testing.T) {
	e := NewFormHidden("hname", "hvalue")
	html := "<input name=\"hname\" type=\"hidden\" value=\"hvalue\" />"

	require.Equal(t, html, Piece(e).String())
}
