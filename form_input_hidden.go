package dhtml

import (
	"github.com/mitoteam/mttools"
)

type HiddenFormItem struct {
	FormItemBase
}

// force interfaces implementation
var _ FormItemI = (*HiddenFormItem)(nil)

func NewFormHidden(name string, value any) *HiddenFormItem {
	fi := &HiddenFormItem{}
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
