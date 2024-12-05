package dhtml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiv(t *testing.T) {
	div := Div().Id("ay_di").Class("cls").Title("title").Attribute("a", "b c").Attribute("d", "").Text("text")
	html := "<div id=\"ay_di\" class=\"cls\" a=\"b c\" d title=\"title\">text</div>"

	require.Equal(t, html, div.String())
}
