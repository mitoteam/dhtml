package dhtml

// simple <img> element
type ImageElement struct {
	tag *Tag
}

// force interfaces implementation
var _ ElementI = (*ImageElement)(nil)

func NewImg(src string) *ImageElement {
	return &ImageElement{tag: NewTag("img").Attribute("src", src)}
}

func (e *ImageElement) Title(title string) *ImageElement {
	e.tag.Title(title)
	return e
}

func (e *ImageElement) Alt(alt string) *ImageElement {
	e.tag.Attribute("alt", alt)
	return e
}

func (e *ImageElement) Width(width string) *ImageElement {
	e.tag.Attribute("width", width)
	return e
}

func (e *ImageElement) Height(height string) *ImageElement {
	e.tag.Attribute("height", height)
	return e
}

func (e *ImageElement) Class(v ...any) *ImageElement {
	e.tag.Class(v...)
	return e
}

func (e *ImageElement) GetTags() TagList {
	return e.tag.GetTags()
}
