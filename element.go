package dhtml

import (
	"html"
)

const (
	tagKindNormal = iota
	tagKindComment
	tagKindContent
)

type Element struct {
	kind       int // tag kind
	tag        string
	attributes map[string]string

	id      string
	classes []string

	children []*Element

	content string //comments and raw content
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

func (e *Element) Content(content string) *Element {
	r := &Element{
		kind:    tagKindContent,
		content: content,
	}

	return e.Append(r)
}

func (e *Element) IsContent() bool {
	return e.kind == tagKindContent
}

// Adds html comment as a child to the element
func (e *Element) Comment(content string) *Element {
	r := &Element{
		kind:    tagKindComment,
		content: content,
	}

	return e.Append(r)
}

func (e *Element) IsComment() bool {
	return e.kind == tagKindComment
}
