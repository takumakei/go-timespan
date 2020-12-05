package timespan_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/takumakei/go-timespan"
)

func ExampleParse() {
	span, err := timespan.Parse("0:00+0900/3s 13:24+0900/1h2m3s 0:00Z/3m")
	if err != nil {
		panic(err)
	}
	fmt.Println(span)
	t0, err := time.Parse(time.RFC3339, "2020-12-05T00:00:00+09:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(span.Contains(t0))
	t1, err := time.Parse(time.RFC3339, "2020-12-05T13:24:00+09:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(span.Contains(t1))
	t2, err := time.Parse(time.RFC3339, "2020-12-05T09:00:00+09:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(span.Contains(t2))
	// output:
	// (00:00+0900/3s or 13:24+0900/1h2m3s or 00:00Z/3m0s)
	// true
	// true
	// true
}

func TestTimeSlot(t *testing.T) {
	tests := []struct {
		S string
		D string
		T string
		W bool
	}{
		{"01:00+0900", "0s", "2020-12-03T01:00:00+09:00", false},

		{"01:00+0900", "1h", "2020-12-03T00:00:00+09:00", false},

		{"01:00+0900", "1h", "2020-12-03T01:00:00+09:00", true},
		{"01:00+0900", "1h", "2020-12-03T01:59:59+09:00", true},
		{"01:00+0900", "1h", "2020-12-03T02:00:00+09:00", false},

		{"23:30+0900", "1h", "2020-12-03T23:30:00+09:00", true},
		{"23:30+0900", "1h", "2020-12-03T00:29:29+09:00", true},
		{"23:30+0900", "1h", "2020-12-03T00:30:00+09:00", false},

		{"00:00+0900", "-1h", "2020-12-03T23:00:00+09:00", true},
		{"00:00+0900", "-1h", "2020-12-03T23:59:59+09:00", true},
		{"00:00+0900", "-1h", "2020-12-03T00:00:00+09:00", false},

		{"23:30+0900", "-1h", "2020-12-03T22:30:00+09:00", true},
		{"23:30+0900", "-1h", "2020-12-03T23:29:29+09:00", true},
		{"23:30+0900", "-1h", "2020-12-03T23:30:00+09:00", false},
	}
	for i, test := range tests {
		s := MustParseTime("15:04-0700", test.S)
		d := MustParseDuration(test.D)
		tm := MustParseTime(time.RFC3339, test.T)
		slot := timespan.NewTimeSlot(s, d)
		r := slot.Contains(tm)
		if r != test.W {
			t.Error(i, slot, test.T, test.W, r)
		}
	}
}

func MustParseTime(layout, s string) time.Time {
	t, err := time.Parse(layout, s)
	if err != nil {
		panic(err)
	}
	return t
}

func MustParseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}
