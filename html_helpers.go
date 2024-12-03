package dhtml

func Div() *Tag {
	return NewTag("div")
}

func Span() *Tag {
	return NewTag("span")
}

func Content(content string) *Tag {
	r := &Tag{
		kind: tagKindText,
		text: content,
	}

	return r
}
