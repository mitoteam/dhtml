package dhtml

type FormElement struct {
	tag *Tag
}

// force interfaces implementation
var _ ElementI = (*FormElement)(nil)

// Html form just to render it
func NewForm() *FormElement {
	return &FormElement{
		tag: NewTag("form"),
	}
}

func (f *FormElement) Class(v ...any) *FormElement {
	f.tag.Class(v...)
	return f
}

func (f *FormElement) Method(method string) *FormElement {
	f.tag.Attribute("method", method)
	return f
}

func (f *FormElement) Append(v ...any) *FormElement {
	for _, v := range v {
		f.tag.Append(v)
	}

	return f
}

func (f *FormElement) GetTags() TagList {
	return f.tag.GetTags()
}
