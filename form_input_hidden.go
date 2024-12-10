package dhtml

import (
	"github.com/mitoteam/mttools"
)

type FormHiddenElement struct {
	FormItemBase
}

// force interfaces implementation
var _ FormItemI = (*FormHiddenElement)(nil)

func NewFormHidden(name string, value any) *FormHiddenElement {
	fi := &FormHiddenElement{}
	fi.name = name
	fi.defaultValue = value

	fi.renderF = func() (out HtmlPiece) {
		out.Append(
			NewTag("input").Attribute("type", "hidden").
				Attribute("name", fi.name).
				Attribute("value", mttools.AnyToString(fi.defaultValue)),
		)

		return out
	}

	return fi
}
