package dhtml

type FormControlElement struct {
	name    string
	label   HtmlPiece
	element ElementI
}

// force interfaces implementation
var _ ElementI = (*FormControlElement)(nil)

func NewFormControl(name string) *FormControlElement {
	return &FormControlElement{
		name: SafeId(name),
	}
}

func (fc *FormControlElement) Label(v any) *FormControlElement {
	fc.label.Append(v)
	return fc
}

func (fc *FormControlElement) Element(e ElementI) *FormControlElement {
	fc.element = e
	return fc
}

func (fi *FormControlElement) GetTags() TagsList {
	root_tag := Div().Id("id_" + fi.name).Class("form-item")

	root_tag.Append(Div().Append(fi.label))
	if fi.element != nil {
		root_tag.Append(Div().Append(fi.element))
	}

	return root_tag.GetTags()
}
