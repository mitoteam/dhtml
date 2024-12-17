package dhtml

import (
	"log"
	"reflect"

	"github.com/mitoteam/mttools"
)

// Element is something that can be turned in list of html tags.
// Very simple elements are: tags itself, html comments or just plain text content.
// It could be much more complex things like Bootstrap's card for example.
// Whole HTML document is element as well (see dhtml.Document helper).
type ElementI interface {
	GetTags() TagList
}

func AnyToElement(v any) ElementI {
	if v, ok := v.(ElementI); ok {
		return v
	}

	if s, ok := mttools.AnyToStringOk(v); ok {
		return Text(s)
	}

	log.Panicf("unsupported type: %s", reflect.ValueOf(v).Type().Name())
	return nil
}
