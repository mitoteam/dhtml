package dhtml

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/mitoteam/mttools"
)

// CSS classes list handling
type Classes struct {
	list []string
}

// force interface implementation
var _ fmt.Stringer = (*Classes)(nil)

func NewClasses(v ...any) (c Classes) {
	c.Add(v...)

	return c
}

func (c *Classes) GetClassList() []string {
	return c.list
}

func (c *Classes) String() string {
	return strings.Join(c.list, " ")
}

func (c *Classes) Count() int {
	return len(c.list)
}

func (c *Classes) Add(v ...any) *Classes {
	for _, v := range v {
		c.list = mttools.UniqueSlice(append(c.list, AnyToClasslist(v)...))
	}
	return c
}

func (c *Classes) Prepend(v any) *Classes {
	c.list = mttools.UniqueSlice(append(AnyToClasslist(v), c.list...))
	return c
}

// Adds class(es) from v if no classes from class_set already added
func (c *Classes) AddFromSet(class_set []string, v ...any) *Classes {
	// if no classes from class_set found in c.list
	if len(mttools.SlicesIntersection(c.list, class_set)) == 0 {
		c.Add(v...)
	}

	return c
}

// CSS-classes string (or something stringable) "parser"
func AnyToClasslist(v any) []string {
	if classes, ok := v.(Classes); ok {
		return classes.list
	}

	var list []string

	s, ok := mttools.AnyToStringOk(v)

	if !ok { //it is not a string
		list, ok = v.([]string) // is it string list?

		if !ok {
			log.Panicf("unsupported type: %s", reflect.ValueOf(v).Type().Name())
		}
	}

	if list == nil {
		if s == "" {
			return []string{} //no string or list found, return empty list
		} else {
			list = strings.Fields(strings.TrimSpace(s))
		}
	}

	for index, value := range list {
		list[index] = SafeClassName(value)
	}

	return list
}
