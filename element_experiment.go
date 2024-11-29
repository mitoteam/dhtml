package dhtml

func BuildExperimentHtml() *Element {
	head := Tag("head").
		Append(
			Tag("link").
				Attribute("href", "/assets/vendor/bootstrap.min.css").
				Attribute("rel", "stylesheet"),
		)

	body := Tag("body")

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
		)

	html := Tag("html").
		Comment("some <escaped> comment").
		Append(head).
		Append(body)

	return html
}
