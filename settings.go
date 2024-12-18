package dhtml

// dhtml settings to be overridden from outside packages

type settingsType struct {
	EmptyLabelRendererF func(label string, span *Tag)
}

var settings *settingsType

func Settings() *settingsType {
	return settings
}

// default implementations
func init() {
	settings = &settingsType{
		EmptyLabelRendererF: func(label string, span *Tag) {
			if label == "" {
				label = "empty"
			}

			span.Append("["+label+"]").Attribute("style", "color: grey;")
		},
	}
}
