package timespan

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	// ErrParseTimeSlot is an error.
	ErrParseTimeSlot = errors.New("can't parse as TimeSlot")

	// ErrParseTime is an error.
	ErrParseTime = fmt.Errorf("%w (time)", ErrParseTimeSlot)

	// ErrParseDuration is an error.
	ErrParseDuration = fmt.Errorf("%w (duration)", ErrParseTimeSlot)
)

var spaces = regexp.MustCompile(`\s+`)

// MustParse calls Parse(s), returns a span.
// In case of err != nil, it calls panic(err).
func MustParse(s string) Span {
	span, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return span
}

// Parse parses s and returns a span.
func Parse(s string) (Span, error) {
	var ss []Span
	for _, e := range spaces.Split(strings.TrimSpace(s), -1) {
		s, err := ParseTimeSlot(e)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}
	return Or(Any(ss)), nil
}

// ParseTimeSlot parses s and returns a TimeSlot
func ParseTimeSlot(s string) (*TimeSlot, error) {
	m := strings.SplitN(s, "/", 2)
	if len(m) != 2 {
		if d, err := time.ParseDuration(s); err != nil {
			return nil, err
		} else if d >= 24*time.Hour {
			return NewTimeSlot(time.Time{}, 24*time.Hour), nil
		}
		return nil, fmt.Errorf("%w: %s", ErrParseTimeSlot, s)
	}
	t, err := parseTime(strings.TrimSpace(m[0]))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrParseTime, m[0])
	}
	d, err := time.ParseDuration(strings.TrimSpace(m[1]))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrParseDuration, m[1])
	}
	return NewTimeSlot(t, d), nil
}

var (
	layoutUTC = []string{
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02T15:04:05.999999999-0700",
		"2006-01-02T15:04:05.999999Z07:00",
		"2006-01-02T15:04:05.999999-0700",
		"2006-01-02T15:04:05.999Z07:00",
		"2006-01-02T15:04:05.999-0700",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04Z07:00",
		"2006-01-02T15:04-0700",
		"2006-01-02T15Z07:00",
		"2006-01-02T15-0700",
	}

	layoutLocal = []string{
		"2006-01-2T15:04:05.000000000",
		"2006-01-2T15:04:05.000000",
		"2006-01-2T15:04:05.000",
		"2006-01-2T15:04:05",
		"2006-01-2T15:04",
		"2006-01-2T15",
	}
)

func parseTime(s string) (t time.Time, err error) {
	s = "2006-01-02T" + s
	for _, layout := range layoutUTC {
		t, err = time.Parse(layout, s)
		if err == nil {
			return
		}
	}
	for _, layout := range layoutLocal {
		t, err = time.ParseInLocation(layout, s, time.Local)
		if err == nil {
			return
		}
	}
	return
}
