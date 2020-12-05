package timespan

import (
	"fmt"
	"strings"
	"time"
)

// Any is a composite span.
type Any []Span

var _ Span = Any(nil)

// Or makes an Any span of ss.
func Or(ss ...Span) Span {
	a := flatAny(make([]Span, 0, len(ss)), Any(ss))
	switch len(a) {
	case 0:
		return Never
	case 1:
		return a[0]
	}
	return Any(a)
}

func flatAny(a []Span, s Span) []Span {
	if v, ok := s.(Any); ok {
		for _, e := range v {
			a = flatAny(a, e)
		}
	} else {
		a = append(a, s)
	}
	return a
}

// Contains returns true in case of any one of spans in s returns true.
func (s Any) Contains(t time.Time) bool {
	for _, e := range s {
		if e.Contains(t) {
			return true
		}
	}
	return false
}

// String returns a string representation of s.
func (s Any) String() string {
	a := make([]string, len(s))
	for i, e := range s {
		a[i] = fmt.Sprint(e)
	}
	return "(" + strings.Join(a, " or ") + ")"
}
