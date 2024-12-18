package dhtml

// simple <label> element
type LabelElement struct {
	tag *Tag
}

// force interfaces implementation
var _ ElementI = (*LabelElement)(nil)

func NewLabel() *LabelElement {
	return &LabelElement{tag: NewTag("label")}
}

func (e *LabelElement) For(targetId string) *LabelElement {
	e.tag.Attribute("for", SafeId(targetId))
	return e
}

// <label> contents
func (e *LabelElement) Append(v ...any) *LabelElement {
	e.tag.Append(v...)
	return e
}

func (e *LabelElement) Class(v ...any) *LabelElement {
	e.tag.Class(v...)
	return e
}

func (e *LabelElement) GetTags() TagList {
	return e.tag.GetTags()
}
