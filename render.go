package dhtml

import (
	"html"
	"strings"

	"github.com/mitoteam/mttools"
)

var inline_preferred_tags = [...]string{
	"i", "b",
}

// Renders element with all the children to HTML
func (e *Element) Render() string {
	return e.rawRender(0)
}

// does real job (with recursion)
func (e *Element) rawRender(level int) string {
	indent := strings.Repeat("  ", level)

	var sb strings.Builder

	if e.IsComment() {
		sb.WriteString(indent + "<!--" + html.EscapeString(e.content) + "-->")

		return sb.String()
	}

	if e.IsContent() {
		sb.WriteString(indent + html.EscapeString(e.content))

		return sb.String()
	}

	//check and set attributes
	if e.id != "" {
		e.attributes["id"] = e.id
	}

	//CSS classes
	if len(e.classes) > 0 {
		e.attributes["class"] = strings.Join(mttools.UniqueSlice(e.classes), " ")
	}

	//prepare raw HTML output
	sb.WriteString("<" + e.tag)

	//render attributes
	for name, value := range e.attributes {
		sb.WriteString(" " + name + "=\"" + html.EscapeString(value) + "\"")
	}

	if len(e.children) == 0 && len(e.content) == 0 {
		//self closing tag
		sb.WriteString("/>")
	} else {
		sb.WriteString(">")

		//go deeper (recursion)
		for _, child := range e.children {
			sb.WriteString(child.rawRender(level + 1))
		}

		//closing tag
		sb.WriteString("</" + e.tag + ">")
	}

	return sb.String()
}
