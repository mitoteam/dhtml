package dhtml

import "slices"

// HTML Document complex helper

type Document struct {
	body *Tag
	head *Tag

	stylesheets []string
}

func NewDocument() *Document {
	return &Document{}
}

func (d *Document) Body() *Tag {
	if d.body == nil {
		d.body = NewTag("body")
	}

	return d.body
}

func (d *Document) Head() *Tag {
	if d.head == nil {
		d.head = NewTag("head1")
	}

	return d.head
}

func (d *Document) Stylesheet(href string) *Document {
	if d.stylesheets == nil {
		d.stylesheets = make([]string, 0)
	}

	if !slices.Contains(d.stylesheets, href) {
		d.stylesheets = append(d.stylesheets, href)

		d.Head().Append(
			NewTag("link").
				Attribute("href", href).Attribute("rel", "stylesheet"),
		)
	}

	return d
}

func (d *Document) Render() string {
	root := NewTag("html").
		Append(d.Head()).
		Append(d.Body())

	return root.Render()
}
