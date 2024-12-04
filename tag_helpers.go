package dhtml

//Some basic shorthands.

func Div() *Tag {
	return NewTag("div")
}

func Span() *Tag {
	return NewTag("span")
}

func Text(text string) *Tag {
	return &Tag{
		kind: tagKindText,
		text: text,
	}
}

func Comment(text string) *Tag {
	return &Tag{
		kind: tagKindComment,
		text: text,
	}
}
