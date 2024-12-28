package dhtml

import (
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

// Function to be passed to Walk() or WalkR() methods
type ElementWalkFunc func(e ElementI, args ...any)

func AnyToElement(v any) ElementI {
	if v, ok := v.(ElementI); ok {
		//already an element
		return v
	}

	if s, ok := mttools.AnyToStringOk(v); ok {
		//strings are simple text elements
		return Text(s)
	}

	mttools.PanicWithSignatureF("unsupported type: %s", reflect.ValueOf(v).Type().Name())
	return nil
}
