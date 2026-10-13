package dhtml

// <script> element
type (
	ScriptElement struct {
		src      string
		loadKind string
	}
)

// force interface implementation declaring fake variable
var _ ElementI = (*ScriptElement)(nil)

func NewScriptElement(src string) *ScriptElement {
	e := &ScriptElement{
		src:      src,
		loadKind: "defer",
	}

	return e
}

func (e *ScriptElement) SetAsync() *ScriptElement {
	e.loadKind = "async"
	return e
}

func (e *ScriptElement) GetTags() TagList {
	t := NewTag("script").
		Attribute("src", e.src).
		Attribute(e.loadKind, "")

	return TagList{t}
}
