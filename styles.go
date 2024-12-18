package dhtml

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/mitoteam/mttools"
)

// styles list handling
type Styles struct {
	propertyValues map[string]string // property => value
}

// force interface implementation
var _ fmt.Stringer = (*Styles)(nil)

// Creates new Styles{} (with possible initial values).
func NewStyles(v ...any) Styles {
	st := Styles{
		propertyValues: make(map[string]string),
	}

	st.Add(v...)

	return st
}

func (s *Styles) String() string {
	properties := slices.Collect(maps.Keys(s.propertyValues))
	slices.Sort(properties)

	var sb strings.Builder

	for _, property := range properties {
		value := s.propertyValues[property]

		if property != "" && value != "" {
			if sb.Len() > 0 {
				sb.WriteString(" ")
			}

			sb.WriteString(property + ": " + value + ";")
		}
	}

	return sb.String()
}

func (st *Styles) Count() int {
	return len(st.propertyValues)
}

func (st *Styles) Get(property string) string {
	if value, ok := st.propertyValues[property]; ok {
		return value
	}

	return ""
}

func (st *Styles) Set(property, value string) *Styles {
	st.propertyValues[property] = value
	return st
}

func (st *Styles) Clear() *Styles {
	st.propertyValues = make(map[string]string)
	return st
}

// string (or stringable) "parser"
func (st *Styles) Add(v ...any) *Styles {
	for _, v := range v {
		if s, ok := mttools.AnyToStringOk(v); ok {
			for _, property := range strings.Split(s, ";") {
				parts := strings.Split(property, ":")

				if len(parts) == 2 {
					st.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
				}
			}
		}
	}

	return st
}
