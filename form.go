package dhtml

type FormElement struct {
	formData *FormData
	body     HtmlPiece
}

// force interfaces implementation
var _ ElementI = (*FormElement)(nil)

// Html form just to render it
func NewForm() *FormElement {
	return &FormElement{}
}

func (f *FormElement) Append(v any) *FormElement {
	if f.formData != nil { //managed form
		if e, ok := v.(FormItemI); ok {
			fd := f.formData

			if value, ok := fd.values[e.GetName()]; ok {
				if len(value) == 1 { //all Request.PostForm values are arrays of at least one element
					e.SetValue(value[0])
				} else {
					e.SetValue(value)
				}
			} else {
				fd.values[e.GetName()] = []string{""} //add empty string to data
			}

			fd.labels[e.GetName()] = e.GetLabel()
		}
	}

	// simple not managed form - just added
	f.body.Append(v)

	return f
}

func (f *FormElement) GetFormData() *FormData {
	return f.formData
}

func (f *FormElement) GetTags() TagsList {
	root_tag := NewTag("form").Attribute("method", "post")

	root_tag.Append(f.body)

	return root_tag.GetTags()
}
