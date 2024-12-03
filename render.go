package dhtml

import (
	"html"
	"maps"
	"slices"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/mitoteam/mttools"
)

var inline_preferred_tags = []string{
	"i", "b", "span",
}

// Renders element with all the children to HTML
func (e *Tag) Render() string {
	var sb strings.Builder

	for _, tag := range e.GetTags() {
		tag.renderTag(0, &sb)
	}

	return sb.String()
}

// does real job (with recursion)
func (e *Tag) renderTag(level int, sb *strings.Builder) {
	var indent string

	if level > 0 {
		indent = "\n" + strings.Repeat("  ", level)
	}

	sb.WriteString(indent)

	if e.IsComment() {
		sb.WriteString("<!--" + html.EscapeString(e.content) + "-->")

		return
	}

	if e.IsContent() {
		sb.WriteString(html.EscapeString(e.content))

		return
	}

	//prepare raw HTML output
	sb.WriteString("<" + e.tag)

	e.renderAttributes(sb)

	if len(e.children) == 0 && len(e.content) == 0 {
		//self closing tag
		sb.WriteString("/>")
	} else {
		sb.WriteString(">")

		//go deeper (recursion)
		var previousElement *Tag
		for _, child := range e.children {
			child_level := level + 1
			if e.IsInline() {
				child_level = 0
			}

			//separate two consecutive content elements with space
			if previousElement != nil && child.kind == tagKindContent && previousElement.kind == tagKindContent {
				sb.WriteString(" ")
			}

			child.renderTag(child_level, sb)
			previousElement = child
		}

		//closing tag
		if !e.IsInline() {
			sb.WriteString(indent)
		}

		sb.WriteString("</" + e.tag + ">")
	}
}

// check, set and render attributes
func (e *Tag) renderAttributes(sb *strings.Builder) {
	attributes := orderedmap.NewOrderedMap[string, string]()

	//check and set attributes
	if e.id != "" {
		attributes.Set("id", e.id)
		delete(e.attributes, "id") //prefer e.id over direct attributes
	}

	//CSS classes
	if len(e.classes) > 0 {
		attributes.Set("class", strings.Join(mttools.UniqueSlice(e.classes), " "))
		delete(e.attributes, "class") //prefer e.class over direct attributes
	}

	//other attributes in alphabetical order
	for _, name := range slices.Sorted(maps.Keys(e.attributes)) {
		attributes.Set(name, e.attributes[name])
	}

	//render attributes
	for name, value := range attributes.Iterator() {
		value = strings.TrimSpace(value)

		sb.WriteString(" " + strings.TrimSpace(name))

		if len(value) > 0 {
			sb.WriteString("=\"" + html.EscapeString(value) + "\"")
		}
	}
}

func (e *Tag) IsInline() bool {
	//content has no children so considered inline
	if e.kind == tagKindContent {
		return true
	}

	//has no not inline children
	for _, child := range e.children {
		if !child.IsInline() || (child.kind == tagKindNormal && !slices.Contains(inline_preferred_tags, child.tag)) {
			return false
		}
	}

	return true
}
