package dhtml

type FormElement struct {
	formData *FormData
	tag      *Tag
}

// force interfaces implementation
var _ ElementI = (*FormElement)(nil)

// Html form just to render it
func NewForm() *FormElement {
	return &FormElement{
		tag: NewTag("form").Attribute("method", "post"),
	}
}

func (f *FormElement) Class(v any) *FormElement {
	f.tag.Class(v)
	return f
}

func (f *FormElement) Append(v any) *FormElement {
	if f.formData != nil { //managed form
		if e, ok := v.(FormItemI); ok {
			fd := f.formData

			if value, ok := fd.values.GetOk(e.GetName()); ok {
				e.SetValue(value)
			} else {
				fd.values.Set(e.GetName(), "") //add empty string to data
			}

			fd.labels.Set(e.GetName(), e.GetLabel())
		}
	}

	// simple not managed form - just added
	f.tag.Append(v)

	return f
}

func (f *FormElement) GetFormData() *FormData {
	return f.formData
}

func (f *FormElement) GetTags() TagsList {
	return f.tag.GetTags()
}
