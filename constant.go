package timespan

import (
	"time"
)

var (
	// Always means an entire span.
	Always = constant(true)

	// Never means an empty span.
	Never = constant(false)
)

type constant bool

var _ Span = Always

// Contains returns s of bool.
func (s constant) Contains(time.Time) bool { return bool(s) }
