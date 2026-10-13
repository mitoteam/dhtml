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
	scripts     []*ScriptElement
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

// Sets page charset. Default is "utf-8". If empty string is set, no charset meta tag will be added to the head.
func (d *HtmlDocument) Charset(charset string) *HtmlDocument {
	d.charset = charset
	return d
}

// Sets page title
func (d *HtmlDocument) Title(title string) *HtmlDocument {
	d.title = title
	return d
}

// Sets page favicon
func (d *HtmlDocument) Icon(icon string) *HtmlDocument {
	d.icon = icon
	return d
}

// Adds a stylesheet link to the document head. If the same href is already added, it will not be added again.
func (d *HtmlDocument) Stylesheet(href string) *HtmlDocument {
	if !slices.Contains(d.stylesheets, href) {
		d.stylesheets = append(d.stylesheets, href)

		d.Head().Append(
			NewTag("link").
				Attribute("href", href).Attribute("rel", "stylesheet"),
		)
	}

	return d
}

// Adds <script> element to <head> with "defer" set.
func (d *HtmlDocument) Script(src string) *ScriptElement {
	s := NewScriptElement(src)

	d.scripts = append(d.scripts, s)

	return s
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

	for _, script := range d.scripts {
		head.Append(script)
	}

	root := NewTag("html").
		Append(head).
		Append(d.Body())

	return TagList{root}
}

func (d *HtmlDocument) String() string {
	return "<!DOCTYPE html>\n" + d.GetTags()[0].String()
}
