package dhtml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiv(t *testing.T) {
	div := Div().Class("cls")
	html := "<div class=\"cls\"></div>"

	require.Equal(t, html, div.String())
}
