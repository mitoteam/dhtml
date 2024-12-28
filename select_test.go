package dhtml_test

import (
	"testing"

	"github.com/mitoteam/dhtml"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	selectElement := dhtml.NewSelect()

	selectElement.Option("a", "AAA")
	selectElement.Option("b", "BBB").Selected(true)

	html := "<select>\n  <option value=\"a\">AAA</option>\n  <option selected value=\"b\">BBB</option></select>"

	require.Equal(t, html, dhtml.Piece(selectElement).String())
}
