package dhtml

import (
	"html"
	"maps"
	"slices"
	"strings"

	"github.com/mitoteam/mttools"
)

var inline_preferred_tags = []string{
	"i", "b", "span",
}

// Renders element with all the children to HTML
func (e *Element) Render() string {
	return e.rawRender(0)
}

// does real job (with recursion)
func (e *Element) rawRender(level int) string {
	var indent string

	if level > 0 {
		indent = "\n" + strings.Repeat("  ", level)
	}

	var sb strings.Builder
	sb.WriteString(indent)

	if e.IsComment() {
		sb.WriteString("<!--" + html.EscapeString(e.content) + "-->")

		return sb.String()
	}

	if e.IsContent() {
		sb.WriteString(html.EscapeString(e.content))

		return sb.String()
	}

	//prepare raw HTML output
	sb.WriteString("<" + e.tag)

	e.renderAttributes(&sb)

	if len(e.children) == 0 && len(e.content) == 0 {
		//self closing tag
		sb.WriteString("/>")
	} else {
		sb.WriteString(">")

		//go deeper (recursion)
		for _, child := range e.children {
			child_level := level + 1
			if e.isInline() {
				child_level = 0
			}

			sb.WriteString(child.rawRender(child_level))
		}

		//closing tag
		if !e.isInline() {
			sb.WriteString(indent)
		}

		sb.WriteString("</" + e.tag + ">")
	}

	return sb.String()
}

// check, set and render attributes
func (e *Element) renderAttributes(sb *strings.Builder) {
	attributes := make(map[string]string)

	//check and set attributes
	if e.id != "" {
		attributes["id"] = e.id
		delete(e.attributes, "id") //prefer e.id over direct attributes
	}

	//CSS classes
	if len(e.classes) > 0 {
		attributes["class"] = strings.Join(mttools.UniqueSlice(e.classes), " ")
		delete(e.attributes, "class") //prefer e.class over direct attributes
	}

	//other attributes
	maps.Copy(attributes, e.attributes)

	//render attributes
	for name, value := range attributes {
		value = strings.TrimSpace(value)

		sb.WriteString(" " + strings.TrimSpace(name))

		if len(value) > 0 {
			sb.WriteString("=\"" + html.EscapeString(value) + "\"")
		}
	}
}

func (e *Element) isInline() bool {
	//content has no children so considered inline
	if e.kind == tagKindContent {
		return true
	}

	//has no not inline children
	for _, child := range e.children {
		if !child.isInline() {
			return false
		}
	}

	return slices.Contains(inline_preferred_tags, e.tag)
}
