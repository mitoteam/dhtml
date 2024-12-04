package dhtml

// HTML Document complex helper

type FormElement struct {
	body HtmlPiece
}

// force interfaces implementation
var _ ElementI = (*FormElement)(nil)

func Form() *FormElement {
	return &FormElement{}
}

func (f FormElement) GetTags() TagsList {
	f.body.Append("Test Form body")

	return f.body.GetTags()
}
