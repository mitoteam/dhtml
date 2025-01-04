package dhtml

// simple <ul> and <ol> elements
type ListElement struct {
	tag   *Tag
	items []*ListItemElement
}

// force interfaces implementation
var _ ElementI = (*ListElement)(nil)

func NewOrderedList() *ListElement {
	e := &ListElement{tag: NewTag("ol")}
	return e
}

func NewUnorderedList() *ListElement {
	e := &ListElement{tag: NewTag("ul")}
	return e
}

func (e *ListElement) Class(v ...any) *ListElement {
	e.tag.Class(v...)
	return e
}

func (e *ListElement) AppendItem(item *ListItemElement) *ListElement {
	e.items = append(e.items, item)
	return e
}

func (e *ListElement) Item(v ...any) *ListItemElement {
	item := NewListItem()
	item.Append(v...)
	e.AppendItem(item)

	return item
}

func (e *ListElement) ItemCount() int {
	return e.tag.ChildrenCount()
}

func (e *ListElement) GetTags() TagList {
	for _, item := range e.items {
		e.tag.Append(item)
	}

	return e.tag.GetTags()
}

// ======================= ListItemElement ==============================

// simple <li> element
type ListItemElement struct {
	tag *Tag
}

// force interfaces implementation
var _ ElementI = (*ListItemElement)(nil)

func NewListItem() *ListItemElement {
	return &ListItemElement{tag: NewTag("li")}
}

func (e *ListItemElement) Class(v ...any) *ListItemElement {
	e.tag.Class(v...)
	return e
}

func (e *ListItemElement) Append(v ...any) *ListItemElement {
	e.tag.Append(v...)
	return e
}

func (e *ListItemElement) GetTags() TagList {
	return e.tag.GetTags()
}
