package dhtml

type (
	ElementI interface {
		GetTags() TagsList
	}

	// helper types
	TagsList []*Tag
)

// region ElementList
type (
	ElementList struct {
		list []ElementI
	}
)

// force interface implementation declaring fake variable
var _ ElementI = (*ElementList)(nil)

// Shorthand helper for NewList() constructor
func Piece[T string | *ElementList](first_element T) *ElementList {
	list := NewElementList()

	//https://ectobit.com/blog/check-type-of-generic-parameter/
	if s, ok := any(first_element).(string); ok {
		list.AppendText(s)
	} else {
		list.AppendList(any(first_element).(*ElementList))
	}

	return list
}

// Actual Constructor
func NewElementList() *ElementList {
	l := &ElementList{
		list: make([]ElementI, 0),
	}

	return l
}

func (l *ElementList) IsEmpty() bool {
	return len(l.list) == 0
}

// Adds single element to list
func (l *ElementList) Append(e ElementI) *ElementList {
	l.list = append(l.list, e)

	return l
}

// Shorthand for Append()
func (l *ElementList) A(e ElementI) *ElementList { return l.Append(e) }

// Adds single element to list
func (l *ElementList) AppendList(another_list *ElementList) *ElementList {
	l.list = append(l.list, another_list.list...)

	return l
}

// Shorthand for AppendList()
func (l *ElementList) AL(another_list *ElementList) *ElementList {
	return l.AppendList(another_list)
}

// Adds text element to list
func (l *ElementList) AppendText(text string) *ElementList {
	l.list = append(l.list, Text(text))

	return l
}

// Shorthand for AppendText()
func (l *ElementList) AT(text string) *ElementList { return l.AppendText(text) }

// ElementI implementation
func (l *ElementList) GetTags() TagsList {
	tag_list := make(TagsList, 0)

	for _, element := range l.list {
		tag_list = append(tag_list, element.GetTags()...)
	}

	return tag_list
}

//endregion
