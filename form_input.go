package dhtml

type FormInputElement struct {
	inputType string
}

// force interfaces implementation
var _ ElementI = (*FormInputElement)(nil)

func NewFormInput(inputType string) *FormInputElement {
	return &FormInputElement{
		inputType: inputType,
	}
}

func (fi *FormInputElement) GetTags() TagsList {
	return NewTag("input").
		Attribute("type", fi.inputType).
		GetTags()
}
