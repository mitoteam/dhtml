package dhtml

import (
	"html"
	"strings"

	"github.com/mitoteam/mttools"
)

type Element struct {
	tag     string
	id      string
	content string

	attributes map[string]string
	classes    []string
	children   []*Element
}

func Tag(tag string) *Element {
	r := &Element{
		tag: SafeTagName(tag),

		attributes: make(map[string]string),
		classes:    make([]string, 0),

		children: make([]*Element, 0),
	}

	return r
}

func (e *Element) Render() string {
	//check and set attributes
	if e.id != "" {
		e.attributes["id"] = e.id
	}

	//CSS classes
	if len(e.classes) > 0 {
		e.attributes["class"] = strings.Join(mttools.UniqueSlice(e.classes), " ")
	}

	//prepare raw HTML output
	var sb strings.Builder
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

		sb.WriteString(html.EscapeString(e.content))

		//go deeper (recursion)
		for _, child := range e.children {
			sb.WriteString(child.Render())
		}

		//closing tag
		sb.WriteString("</" + e.tag + ">")
	}

	return sb.String()
}

// Adds child element
func (e *Element) Append(child_element *Element) *Element {
	e.children = append(e.children, child_element)
	return e
}

// Adds child element to the beginning of children list
func (e *Element) Prepend(child_element *Element) *Element {
	e.children = append([]*Element{child_element}, e.children...)
	return e
}

func (e *Element) Content(content string) *Element {
	e.content = content
	return e
}

func (e *Element) Id(id string) *Element {
	e.id = id
	return e
}

func (e *Element) Attribute(name, value string) *Element {
	e.attributes[SafeAttributeName(name)] = html.EscapeString(value)
	return e
}

func (e *Element) GetAttribute(name string) string {
	return e.attributes[name]
}

func (e *Element) Class(name string) *Element {
	e.classes = append(e.classes, SafeClassName(name))
	return e
}
