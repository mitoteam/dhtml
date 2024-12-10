package dhtml

import (
	"fmt"
	"slices"
)

// HTML Document complex helper

type HtmlDocument struct {
	body *Tag
	head *Tag

	stylesheets []string
}

// force interfaces implementation
var _ fmt.Stringer = (*HtmlDocument)(nil)
var _ ElementI = (*HtmlDocument)(nil)

func NewHtmlDocument() *HtmlDocument {
	return &HtmlDocument{}
}

func (d *HtmlDocument) Body() *Tag {
	if d.body == nil {
		d.body = NewTag("body")
	}

	return d.body
}

func (d *HtmlDocument) Head() *Tag {
	if d.head == nil {
		d.head = NewTag("head")
	}

	return d.head
}

func (d *HtmlDocument) Title(title string) *HtmlDocument {
	d.Head().Append(NewTag("title").Text(title))
	return d
}

func (d *HtmlDocument) Stylesheet(href string) *HtmlDocument {
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

func (d *HtmlDocument) GetTags() TagsList {
	root := NewTag("html").
		Append(d.Head()).
		Append(d.Body())

	return TagsList{root}
}

func (d *HtmlDocument) String() string {
	return d.GetTags()[0].String()
}
