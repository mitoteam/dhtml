package dhtml

type FormSubmitElement struct {
	FormItemExtBase
	classes Classes
}

// force interfaces implementation
var _ FormItemExtI = (*FormSubmitElement)(nil)

func NewFormSubmit() *FormSubmitElement {
	fi := &FormSubmitElement{}
	fi.wrapped = true
	fi.name = "submit"
	fi.classes = settings.DefaultSubmitButtonClasses

	fi.renderF = func() (out HtmlPiece) {
		if fi.label.IsEmpty() {
			fi.label.AppendText("Submit")
		}

		out.Append(
			NewTag("button").Attribute("type", "submit").
				Append(fi.label).
				Class(fi.classes),
		)

		return out
	}

	return fi
}

func (fi *FormSubmitElement) Label(v any) *FormSubmitElement {
	fi.label.Append(v)
	return fi
}

func (fi *FormSubmitElement) Class(v any) *FormSubmitElement {
	fi.classes.Add(v)
	return fi
}
