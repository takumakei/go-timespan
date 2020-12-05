package timespan

import "time"

// Span represents a time span.
type Span interface {
	// Contains returns true if t is inside a span of this.
	Contains(t time.Time) bool
}
