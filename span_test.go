package timespan_test

import (
	"testing"
	"time"

	"github.com/takumakei/go-timespan"
)

func TestSpan(t *testing.T) {
	tests := []struct {
		S string
		T time.Time
		W bool
	}{
		{"00:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T09:00:00+09:00").Add(-1), false},
		{"00:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T09:00:00+09:00"), true},
		{"00:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T09:00:00+09:00").Add(time.Hour).Add(-1), true},
		{"00:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T09:00:00+09:00").Add(time.Hour), false},

		{"15:30Z/1h", MustParseTime(time.RFC3339, "2020-12-05T00:30:00+09:00").Add(-1), false},
		{"15:30Z/1h", MustParseTime(time.RFC3339, "2020-12-05T00:30:00+09:00"), true},
		{"15:30Z/1h", MustParseTime(time.RFC3339, "2020-12-05T00:30:00+09:00").Add(time.Hour).Add(-1), true},
		{"15:30Z/1h", MustParseTime(time.RFC3339, "2020-12-05T00:30:00+09:00").Add(time.Hour), false},

		{"23:30Z/1h", MustParseTime(time.RFC3339, "2020-12-05T08:30:00+09:00").Add(-1), false},
		{"23:30Z/1h", MustParseTime(time.RFC3339, "2020-12-05T08:30:00+09:00"), true},
		{"23:30Z/1h", MustParseTime(time.RFC3339, "2020-12-05T08:30:00+09:00").Add(time.Hour).Add(-1), true},
		{"23:30Z/1h", MustParseTime(time.RFC3339, "2020-12-05T08:30:00+09:00").Add(time.Hour), false},

		{"00:00+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T00:00:00+09:00").Add(-1), false},
		{"00:00+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T00:00:00+09:00"), true},
		{"00:00+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T00:00:00+09:00").Add(time.Hour).Add(-1), true},
		{"00:00+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T00:00:00+09:00").Add(time.Hour), false},

		{"15:30+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T15:30:00+09:00").Add(-1), false},
		{"15:30+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T15:30:00+09:00"), true},
		{"15:30+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T15:30:00+09:00").Add(time.Hour).Add(-1), true},
		{"15:30+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T15:30:00+09:00").Add(time.Hour), false},

		{"23:30+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T23:30:00+09:00").Add(-1), false},
		{"23:30+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T23:30:00+09:00"), true},
		{"23:30+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T23:30:00+09:00").Add(time.Hour).Add(-1), true},
		{"23:30+0900/1h", MustParseTime(time.RFC3339, "2020-12-05T23:30:00+09:00").Add(time.Hour), false},

		{"00:00Z/1h 12:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T09:00:00+09:00").Add(-1), false},
		{"00:00Z/1h 12:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T09:00:00+09:00"), true},
		{"00:00Z/1h 12:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T09:00:00+09:00").Add(time.Hour).Add(-1), true},
		{"00:00Z/1h 12:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T09:00:00+09:00").Add(time.Hour), false},

		{"00:00Z/1h 12:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T21:00:00+09:00").Add(-1), false},
		{"00:00Z/1h 12:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T21:00:00+09:00"), true},
		{"00:00Z/1h 12:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T21:00:00+09:00").Add(time.Hour).Add(-1), true},
		{"00:00Z/1h 12:00Z/1h", MustParseTime(time.RFC3339, "2020-12-05T21:00:00+09:00").Add(time.Hour), false},
	}
	for i, test := range tests {
		s, err := timespan.Parse(test.S)
		if err != nil {
			t.Error(i, test.S, err)
		}
		r := s.Contains(test.T)
		if r != test.W {
			t.Error(i, test.S, test.T, r, test.W)
		}
	}
}
