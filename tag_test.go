package dhtml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiv(t *testing.T) {
	div := Div().Id("ay_di").Title("title").
		Class("cls").Class("with space").Class("cls" /* duplicate */, "another-one").
		Attribute("a", "b c").Attribute("d", "").Text("text")
	html := "<div id=\"ay_di\" class=\"cls with space another-one\" a=\"b c\" d title=\"title\">text</div>"

	require.Equal(t, html, div.String())
}

func TestStyles(t *testing.T) {
	div := Div().Class("cls").Text("text").
		Style("border", "1px solid grey").
		Styles("text-align    :     right;    color: blue;")

	html := "<div class=\"cls\" style=\"border: 1px solid grey; color: blue; text-align: right;\">text</div>"

	require.Equal(t, html, div.String())
}
