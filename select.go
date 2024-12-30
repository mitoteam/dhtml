package dhtml

import "github.com/mitoteam/mttools"

type SelectElement struct {
	tag     *Tag
	options []*OptionElement
}

// force interface implementation declaring fake variable
var _ ElementI = (*SelectElement)(nil)

func NewSelect() *SelectElement {
	return &SelectElement{
		tag: NewTag("select"),
	}
}

func (c *SelectElement) Class(v any) *SelectElement {
	c.tag.Class(v)
	return c
}

func (c *SelectElement) Id(id string) *SelectElement {
	c.tag.Id(id)
	return c
}

func (c *SelectElement) Attribute(name, value string) *SelectElement {
	c.tag.Attribute(name, value)
	return c
}

func (c *SelectElement) Option(value any, body any) *OptionElement {
	o := NewOption().Value(value).Body(body)

	c.AppendOption(o)

	return o
}

func (c *SelectElement) AppendOption(option *OptionElement) *SelectElement {
	c.options = append(c.options, option)
	return c
}

func (c *SelectElement) GetTags() TagList {
	for _, option := range c.options {
		c.tag.Append(option)
	}

	return c.tag.GetTags()
}

// ==================== OptionElement ===================
type OptionElement struct {
	value    string
	selected bool
	body     HtmlPiece
}

// force interface implementation declaring fake variable
var _ ElementI = (*OptionElement)(nil)

func NewOption() *OptionElement {
	return &OptionElement{}
}

func (c *OptionElement) Value(v any) *OptionElement {
	c.value = mttools.AnyToString(v)
	return c
}

func (c *OptionElement) Body(v any) *OptionElement {
	c.body.Append(v)
	return c
}

func (c *OptionElement) Selected(b bool) *OptionElement {
	c.selected = b
	return c
}

// region Rendering
func (c *OptionElement) GetTags() TagList {
	selectTag := NewTag("option").Append(c.body)

	selectTag.Attribute("value", c.value)

	if c.selected {
		selectTag.Attribute("selected", "")
	}

	return selectTag.GetTags()
}
