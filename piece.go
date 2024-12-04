package dhtml

import (
	"fmt"
	"log"
	"reflect"
)

// HtmlPiece is set of one or several html elements (or no elements at all). Could be tags, complex elements, text content etc.
// Every HtmlPiece as an element itself.
type HtmlPiece struct {
	list []ElementI
}

// force interface implementation declaring fake variable
var _ ElementI = (*HtmlPiece)(nil)

// Shorthand helper for NewHtmlPiece() constructor
func Piece(first_element any) *HtmlPiece {
	return NewHtmlPiece().AppendElement(AnyToElement(first_element))
}

// Actual Constructor
func NewHtmlPiece() *HtmlPiece {
	l := &HtmlPiece{
		list: make([]ElementI, 0),
	}

	return l
}

func (l *HtmlPiece) IsEmpty() bool {
	return len(l.list) == 0
}

// Adds something to list
func (l *HtmlPiece) Append(v any) *HtmlPiece {
	switch v := v.(type) {
	case *HtmlPiece:
		return l.AppendList(v)

	case ElementI:
		return l.AppendElement(v)

	default:
		return l.AppendElement(AnyToElement(v))
	}
}

// Adds single element to list
func (l *HtmlPiece) AppendElement(e ElementI) *HtmlPiece {
	l.list = append(l.list, e)

	return l
}

// Shorthand for Append()
func (l *HtmlPiece) AE(e ElementI) *HtmlPiece { return l.AppendElement(e) }

// Adds single element to list
func (l *HtmlPiece) AppendList(another_list *HtmlPiece) *HtmlPiece {
	l.list = append(l.list, another_list.list...)

	return l
}

// Shorthand for AppendList()
func (l *HtmlPiece) AL(another_list *HtmlPiece) *HtmlPiece {
	return l.AppendList(another_list)
}

// Adds text element to list
func (l *HtmlPiece) AppendText(text string) *HtmlPiece {
	l.list = append(l.list, Text(text))

	return l
}

// Shorthand for AppendText()
func (l *HtmlPiece) AT(text string) *HtmlPiece { return l.AppendText(text) }

// ElementI implementation
func (l *HtmlPiece) GetTags() TagsList {
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
