package dhtml

import (
	"fmt"
	"log"
	"reflect"
)

type (
	ElementI interface {
		GetTags() TagsList
	}

	// helper types
	TagsList []*Tag
)

//region ElementList

type ElementList struct {
	list []ElementI
}

// force interface implementation declaring fake variable
var _ ElementI = (*ElementList)(nil)

// Shorthand helper for NewElementList() constructor
func Piece(first_element any) *ElementList {
	return NewElementList().AppendElement(AnyToElement(first_element))
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

// Adds something to list
func (l *ElementList) Append(v any) *ElementList {
	switch v := v.(type) {
	case *ElementList:
		return l.AppendList(v)

	case ElementI:
		return l.AppendElement(v)

	default:
		return l.AppendElement(AnyToElement(v))
	}
}

// Adds single element to list
func (l *ElementList) AppendElement(e ElementI) *ElementList {
	l.list = append(l.list, e)

	return l
}

// Shorthand for Append()
func (l *ElementList) AE(e ElementI) *ElementList { return l.AppendElement(e) }

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

//region Helper functions

// Helper to convert any value to ElementI
func AnyToElement(v any) ElementI {
	//https://stackoverflow.com/questions/72267243/unioning-an-interface-with-a-type-in-golang
	switch v := v.(type) {
	case ElementI:
		return v

	case string:
		return Text(v)

	case fmt.Stringer:
		return Text(v.String())
	}

	// handle the remaining type set of ~string
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.String {
		return Text(r.String())
	}

	log.Panicf("unsupported type: %s", r.Type().Name())
	return nil
}

//endregion
