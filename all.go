package timespan

import (
	"fmt"
	"strings"
	"time"
)

// All is a composite span.
type All []Span

var _ Span = All(nil)

// And makes an All span of ss.
func And(ss ...Span) Span {
	a := flatAll(make([]Span, 0, len(ss)), All(ss))
	switch len(a) {
	case 0:
		return Never
	case 1:
		return a[0]
	}
	return All(a)
}

func flatAll(a []Span, s Span) []Span {
	if v, ok := s.(All); ok {
		for _, e := range v {
			a = flatAll(a, e)
		}
	} else {
		a = append(a, s)
	}
	return a
}

// Contains returns true in case of all of spans in s returns true.
func (s All) Contains(t time.Time) bool {
	for _, e := range s {
		if !e.Contains(t) {
			return false
		}
	}
	return true
}

// String returns a string representation of s.
func (s All) String() string {
	a := make([]string, len(s))
	for i, e := range s {
		a[i] = fmt.Sprint(e)
	}
	return "(" + strings.Join(a, " and ") + ")"
}
