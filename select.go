package dhtml

import "github.com/mitoteam/mttools"

type SelectElement struct {
	classes Classes
	options []*OptionElement
}

// force interface implementation declaring fake variable
var _ ElementI = (*SelectElement)(nil)

func NewSelect() *SelectElement {
	return &SelectElement{}
}

func (c *SelectElement) Class(v any) *SelectElement {
	c.classes.Add(v)
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
	selectTag := NewTag("select").Class(c.classes)

	for _, option := range c.options {
		selectTag.Append(option)
	}

	return selectTag.GetTags()
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
