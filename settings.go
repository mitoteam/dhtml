package dhtml

// dhtml settings to be overridden from outside packages

type settingsType struct {
	//TODO: SubmitElementRendererF func() FormSubmitElement
	DefaultSubmitButtonClasses Classes

	// function to render errors block if there are any after form validation
	FormErrorsRendererF func(fd *FormData) (out HtmlPiece)

	EmptyLabelRendererF func(label string, span *Tag)
}

var settings *settingsType

func Settings() *settingsType {
	return settings
}

// default implementations
func init() {
	settings = &settingsType{
		FormErrorsRendererF: func(fd *FormData) (out HtmlPiece) {
			container := Div().Class("form-errors")

			for name, itemErrors := range fd.errorList {
				for _, itemError := range itemErrors {
					errorOut := Div().Class("item-error")

					if name != "" {
						errorOut.Attribute("data-form-item-name", name).
							Append(Span().Append(fd.GetLabel(name))).
							Text(":")
					}

					errorOut.Append(itemError)

					container.Append(errorOut)
				}
			}

			out.Append(container)
			return out
		},

		EmptyLabelRendererF: func(label string, span *Tag) {
			if label == "" {
				label = "empty"
			}

			span.Append("["+label+"]").Attribute("style", "color: grey;")
		},
	}
}
