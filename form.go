package dhtml

type FormElement struct {
	body HtmlPiece
}

// force interfaces implementation
var _ ElementI = (*FormElement)(nil)

func NewForm() *FormElement {
	return &FormElement{}
}

func (f *FormElement) Append(v any) *FormElement {
	f.body.Append(v)
	return f
}

func (f *FormElement) GetTags() TagsList {
	root_tag := NewTag("form").Attribute("method", "post")

	root_tag.Append(f.body)

	return root_tag.GetTags()
}
