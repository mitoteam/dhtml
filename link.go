package dhtml

// simple <a> element
type LinkElement struct {
	tag *Tag
}

// force interfaces implementation
var _ ElementI = (*LinkElement)(nil)

// Html form just to render it
func NewLink(href string) *LinkElement {
	return &LinkElement{tag: NewTag("a").Attribute("href", href)}
}

func (e *LinkElement) Target(target string) *LinkElement {
	e.tag.Attribute("target", target)
	return e
}

func (e *LinkElement) Label(v any) *LinkElement {
	e.tag.Append(v)
	return e
}

func (e *LinkElement) Title(title string) *LinkElement {
	e.tag.Title(title)
	return e
}

func (e *LinkElement) Class(v any) *LinkElement {
	e.tag.Class(v)
	return e
}

func (e *LinkElement) GetTags() TagsList {
	return e.tag.GetTags()
}
