package timespan

import (
	"time"
)

// TimeSlot is a Span represents some hours out of 24 hours.
//
// Begin time is inclusive. End time is exclusive.
// In case of d > 0, t is of begin time, end time is t.Add(d).
// In case of d < 0, t is of end time, begin time is t.Add(d).
type TimeSlot struct {
	t time.Time
	d time.Duration

	a time.Duration
	b time.Duration
	f func(time.Time) bool
}

// NewTimeSlot returns a TimeSlot consisted by reference time and duration.
func NewTimeSlot(t time.Time, d time.Duration) *TimeSlot {
	u := t.UTC()
	v := d
	if v < 0 {
		u = u.Add(v)
		v = -v
	}
	a := offset(u)
	b := offset(u.Add(v))
	s := &TimeSlot{
		t: t,
		d: d,
		a: a,
		b: b,
	}
	if a <= b {
		s.f = s.and
	} else {
		s.f = s.or
	}
	return s
}

func offset(t time.Time) time.Duration {
	return t.Sub(t.Truncate(24 * time.Hour))
}

// String returns a string representation of s.
func (s *TimeSlot) String() string {
	if s.d >= 24*time.Hour {
		return s.d.String()
	}
	layout := "15:04Z0700"
	switch {
	case s.t.Nanosecond() != 0:
		layout = "15:04:05.999999999Z0700"
	case s.t.Second() != 0:
		layout = "15:04Z0700"
	}
	return s.t.Format(layout) + "/" + s.d.String()
}

// Contains returns true if t is inside a span of s.
func (s *TimeSlot) Contains(t time.Time) bool {
	return s.f(t)
}

func (s *TimeSlot) and(t time.Time) bool {
	x := offset(t.UTC())
	return s.a <= x && x < s.b
}

func (s *TimeSlot) or(t time.Time) bool {
	x := offset(t.UTC())
	return s.a <= x || x < s.b
}
