package dhtml

import (
	"html"
)

const (
	tagKindNormal = iota
	tagKindComment
	tagKindContent
)

type (
	// Basic tag element implementation
	Tag struct {
		kind       int // tag kind
		tag        string
		attributes map[string]string

		id      string
		classes []string

		children []*Tag

		content string //comments and raw content
	}
)

// forcing interface implementation
var _ ElementI = &Tag{}

func NewTag(tag string) *Tag {
	r := &Tag{
		tag: SafeTagName(tag),

		attributes: make(map[string]string),
		classes:    make([]string, 0),

		children: make([]*Tag, 0),
	}

	return r
}

func (e *Tag) GetTags() []*Tag {
	//tag is a list of itself
	return []*Tag{e}
}

// Adds child element
func (e *Tag) Append(child_element *Tag) *Tag {
	e.children = append(e.children, child_element)
	return e
}

// Adds child element to the beginning of children list
func (e *Tag) Prepend(child_element *Tag) *Tag {
	e.children = append([]*Tag{child_element}, e.children...)
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

func (e *Tag) Content(content string) *Tag {
	r := &Tag{
		kind:    tagKindContent,
		content: content,
	}

	return e.Append(r)
}

func (e *Tag) IsContent() bool {
	return e.kind == tagKindContent
}

// Adds html comment as a child to the element
func (e *Tag) Comment(content string) *Tag {
	r := &Tag{
		kind:    tagKindComment,
		content: content,
	}

	return e.Append(r)
}

func (e *Tag) IsComment() bool {
	return e.kind == tagKindComment
}
