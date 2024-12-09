package dhtml

import (
	"fmt"
	"html"
	"log"
	"maps"
	"slices"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
)

const (
	tagKindNormal = iota
	tagKindComment
	tagKindText
)

var (
	inline_preferred_tags = []string{
		"i", "em", "b", "strong", "span", "small", "del", "ins",
	}

	//https://html.spec.whatwg.org/multipage/syntax.html#void-elements
	void_tags = []string{
		"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "source", "track", "wbr",
	}
)

type (
	// Basic tag element implementation
	Tag struct {
		kind       int // tag kind
		tag        string
		attributes map[string]string

		id      string
		classes Classes

		children HtmlPiece

		text string //comments and raw text content
	}
)

// force interfaces implementation
var _ ElementI = (*Tag)(nil)
var _ fmt.Stringer = (*Tag)(nil)

// Tag constructor
func NewTag(tag string) *Tag {
	r := &Tag{
		tag: SafeTagName(tag),

		attributes: make(map[string]string),
	}

	return r
}

func (t *Tag) GetTags() TagsList {
	//tag is a list of itself
	return TagsList{t}
}

// Adds child element
func (t *Tag) Append(v any) *Tag {
	t.children.Append(v)
	return t
}

func (e *Tag) Id(id string) *Tag {
	e.id = SafeId(id)
	return e
}

// Sets attribute.
func (e *Tag) Attribute(name, value string) *Tag {
	e.attributes[SafeAttributeName(name)] = html.EscapeString(value)
	return e
}

func (e *Tag) GetAttribute(name string) string {
	return e.attributes[name]
}

func (e *Tag) Title(s string) *Tag {
	if s != "" {
		e.Attribute("title", s)
	}
	return e
}

// Adds one or more CSS classes.
func (e *Tag) Class(v any) *Tag {
	e.classes.Add(v)
	return e
}

func (e *Tag) GetClasses() *Classes {
	return &e.classes
}

func (e *Tag) Text(content string) *Tag {
	return e.Append(Text(content))
}

func (e *Tag) Textf(format string, a ...any) *Tag {
	return e.Append(Text(fmt.Sprintf(format, a...)))
}

// Adds html comment as a child to the element
func (e *Tag) Comment(text string) *Tag {
	return e.Append(Comment(text))
}

func (e *Tag) IsText() bool {
	return e.kind == tagKindText
}

func (e *Tag) IsComment() bool {
	return e.kind == tagKindComment
}

func (e *Tag) HasChildren() bool {
	return !e.children.IsEmpty()
}

// true if this tag could be rendered inline, false - should be rendered on new line and indented.
func (t *Tag) IsInline() bool {
	//content or comments has no children so considered inline
	if t.kind == tagKindText || t.kind == tagKindComment {
		return true
	}

	// too many children
	if t.children.GetElementsCount() > 4 {
		return false
	}

	//has no not inline children
	for _, child_tag := range t.children.GetTags() {
		if !child_tag.IsInline() || (child_tag.kind == tagKindNormal && !slices.Contains(inline_preferred_tags, child_tag.tag)) {
			return false
		}
	}

	return true
}

//region Renderer

// Renders element with all the children as HTML
func (t *Tag) String() string {
	var sb strings.Builder

	t.renderTag(0, &sb)

	return sb.String()
}

// does real job (with recursion)
func (t *Tag) renderTag(level int, sb *strings.Builder) {
	var indent string

	if level > 0 {
		indent = "\n" + strings.Repeat("  ", level)
	}

	sb.WriteString(indent)

	if t.IsComment() {
		sb.WriteString("<!--" + html.EscapeString(t.text) + "-->")
		return
	}

	if t.IsText() {
		sb.WriteString(html.EscapeString(t.text))
		return
	}

	//prepare raw HTML output
	sb.WriteString("<" + t.tag)

	t.renderAttributes(sb)

	if slices.Contains(void_tags, t.tag) {
		// void tag
		if t.children.IsEmpty() {
			sb.WriteString(" />")
		} else {
			log.Fatalf("Void tag <%s> can not have children", t.tag)
		}
	} else {
		sb.WriteString(">")

		previousIsContent := false

		//go deeper (recursion)
		child_level := level + 1
		for _, child_tag := range t.children.GetTags() {
			if t.IsInline() {
				child_level = 0
			}

			//separate two consecutive content elements with space
			if previousIsContent && child_tag.kind == tagKindText {
				sb.WriteString(" ")
			}

			child_tag.renderTag(child_level, sb)

			previousIsContent = child_tag.kind == tagKindText
		}

		//closing tag
		if !t.IsInline() {
			sb.WriteString(indent)
		}

		sb.WriteString("</" + t.tag + ">")
	}
}

// check, set and render attributes
func (t *Tag) renderAttributes(sb *strings.Builder) {
	attributes := orderedmap.NewOrderedMap[string, string]()

	//check and set attributes
	if t.id != "" {
		attributes.Set("id", t.id)
		delete(t.attributes, "id") //prefer e.id over direct attributes
	}

	//CSS classes
	if t.classes.GetCount() > 0 {
		attributes.Set("class", t.classes.String())
		delete(t.attributes, "class") //prefer e.class over direct attributes
	}

	//other attributes in alphabetical order
	for _, name := range slices.Sorted(maps.Keys(t.attributes)) {
		attributes.Set(name, t.attributes[name])
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

//endregion
