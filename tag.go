package dhtml

import (
	"html"
	"log"
	"maps"
	"slices"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/mitoteam/mttools"
)

const (
	tagKindNormal = iota
	tagKindComment
	tagKindText
)

var (
	inline_preferred_tags = []string{
		"i", "b", "span",
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
		classes []string

		children ElementsList

		text string //comments and raw text content
	}
)

// force interface implementation
var _ ElementI = &Tag{}

// Tag constructor
func NewTag(tag string) *Tag {
	r := &Tag{
		tag: SafeTagName(tag),

		attributes: make(map[string]string),
		classes:    make([]string, 0),

		children: make(ElementsList, 0),
	}

	return r
}

func (t *Tag) GetTags() TagsList {
	//tag is a list of itself
	return TagsList{t}
}

// Adds child element
func (t *Tag) Append(element ElementI) *Tag {
	t.children = append(t.children, element)
	return t
}

// Adds child element
func (t *Tag) AppendList(list ElementsList) *Tag {
	t.children = append(t.children, list...)
	return t
}

// Adds child element to the beginning of children list
func (e *Tag) Prepend(element ElementI) *Tag {
	e.children = append(ElementsList{element}, e.children...)
	return e
}

func (e *Tag) Id(id string) *Tag {
	e.id = id
	return e
}

func (e *Tag) Attribute(name, value string) *Tag {
	e.attributes[SafeAttributeName(name)] = html.EscapeString(value)
	return e
}

func (e *Tag) GetAttribute(name string) string {
	return e.attributes[name]
}

func (e *Tag) Class(class_name string) *Tag {
	e.classes = append(e.classes, SafeClassName(class_name))
	return e
}

func (e *Tag) Classes(classes []string) *Tag {
	for _, class_name := range classes {
		SafeClassName(class_name)
	}

	e.classes = append(e.classes, classes...)
	return e
}

func (e *Tag) Text(content string) *Tag {
	r := &Tag{
		kind: tagKindText,
		text: content,
	}

	return e.Append(r)
}

func (e *Tag) IsText() bool {
	return e.kind == tagKindText
}

// Adds html comment as a child to the element
func (e *Tag) Comment(content string) *Tag {
	r := &Tag{
		kind: tagKindComment,
		text: content,
	}

	return e.Append(r)
}

func (e *Tag) IsComment() bool {
	return e.kind == tagKindComment
}

// true if this tag could be rendered inline, false - should be rendered on new line and indented.
func (t *Tag) IsInline() bool {
	//content has no children so considered inline
	if t.kind == tagKindText {
		return true
	}

	//has no not inline children
	for _, child := range t.children {
		if tag, ok := child.(*Tag); ok {
			if !tag.IsInline() || (tag.kind == tagKindNormal && !slices.Contains(inline_preferred_tags, tag.tag)) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// #region Renderer
// Renders element with all the children as HTML
func (t *Tag) Render() string {
	var sb strings.Builder

	for _, tag := range t.GetTags() {
		tag.renderTag(0, &sb)
	}

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
		if len(t.children) == 0 {
			sb.WriteString("/>")
		} else {
			log.Fatalf("Void tag <%s> can not have children", t.tag)
		}
	} else {
		sb.WriteString(">")

		previousIsContent := false

		//go deeper (recursion)
		for _, child_element := range t.children {
			child_level := level + 1

			for _, child_tag := range child_element.GetTags() {
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
	if len(t.classes) > 0 {
		attributes.Set("class", strings.Join(mttools.UniqueSlice(t.classes), " "))
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

//#endregion
