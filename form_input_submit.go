package dhtml

type SubmitFormItem struct {
	FormItemExtBase
	classes Classes
}

// force interfaces implementation
var _ FormItemExtI = (*SubmitFormItem)(nil)

var defaultSubmitButtonClasses Classes

func SetDefaultSubmitButtonClasses(v any) {
	defaultSubmitButtonClasses = AnyToClasses(v)
}

func NewFormSubmit() *SubmitFormItem {
	fi := &SubmitFormItem{}
	fi.name = "submit"
	fi.classes = defaultSubmitButtonClasses

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

func (fi *SubmitFormItem) Label(v any) *SubmitFormItem {
	fi.label.Append(v)
	return fi
}

func (fi *SubmitFormItem) Class(v any) *SubmitFormItem {
	fi.classes.Add(v)
	return fi
}
