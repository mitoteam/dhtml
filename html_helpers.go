package dhtml

func Div() *Tag {
	return NewTag("div")
}

func Span() *Tag {
	return NewTag("span")
}

func Text(text string) *Tag {
	r := &Tag{
		kind: tagKindText,
		text: text,
	}

	return r
}
