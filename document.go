package dhtml

import (
	"fmt"
	"slices"
)

// HTML Document complex helper

type HtmlDocument struct {
	body *Tag
	head *Tag

	//metadata for head
	stylesheets []string
	charset     string
	title       string
	icon        string
}

// force interfaces implementation
var _ fmt.Stringer = (*HtmlDocument)(nil)
var _ ElementI = (*HtmlDocument)(nil)

func NewHtmlDocument() *HtmlDocument {
	return &HtmlDocument{
		head:    NewTag("head"),
		body:    NewTag("body"),
		charset: "utf-8",
	}
}

func (d *HtmlDocument) Body() *Tag {
	return d.body
}

func (d *HtmlDocument) Head() *Tag {
	return d.head
}

func (d *HtmlDocument) Charset(charset string) *HtmlDocument {
	d.charset = charset
	return d
}

func (d *HtmlDocument) Title(title string) *HtmlDocument {
	d.title = title
	return d
}

func (d *HtmlDocument) Icon(icon string) *HtmlDocument {
	d.icon = icon
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

func (d *HtmlDocument) GetTags() TagList {
	head := d.Head()

	if d.charset != "" {
		head.Append(NewTag("meta").Attribute("charset", d.charset))
	}

	if d.title != "" {
		head.Append(NewTag("title").Text(d.title))
	}

	if d.icon != "" {
		head.Append(NewTag("link").Attribute("rel", "icon").Attribute("href", d.icon))
	}

	root := NewTag("html").
		Append(head).
		Append(d.Body())

	return TagList{root}
}

func (d *HtmlDocument) String() string {
	return "<!DOCTYPE html>\n" + d.GetTags()[0].String()
}
