package dhtml

func BuildExperimentHtml() *Tag {
	head := NewTag("head").
		Append(
			NewTag("link").
				Attribute("href", "/assets/vendor/bootstrap.min.css").
				Attribute("rel", "stylesheet"),
		)

	body := NewTag("body")

	div := Div().
		Id("test").
		Attribute("data-mt-test", "some attribute").
		//classes deduplication
		Class("border").Class("m-3").Class("p-3").Class("border").
		Content("some <escaped> text")

	body.Append(div)

	body.
		Append(
			Div().Class("border").Class("mt-3").Class("p-3").
				Content("multi").
				Append(
					Span().Content("red").Classes([]string{"border", "border-danger", "border-5"}),
				).
				Content("content"),
		).
		Append(
			Div().Classes([]string{"border", "p-3", "m-3"}).
				Content("content").
				Content("only"),
		)

	html := NewTag("html").
		Comment("some <escaped> comment").
		Append(head).
		Append(body)

	return html
}
