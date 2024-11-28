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
		Class("border").Class("m-3").Class("p-3").Class("border").
		Content("some text")
	body.Append(div)

	html := Tag("html").
		Append(head).
		Append(body)

	return html
}
