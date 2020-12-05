package timespan_test

import (
	"fmt"
	"testing"
	"time"
	_ "time/tzdata"

	"github.com/takumakei/go-timespan"
)

func Example() {
	// the span represents 2 hours from 23:00Z to 01:00Z(next day)
	span := timespan.MustParse("23:00Z/2h")
	epoch := time.Unix(0, 0)
	fmt.Println(span.Contains(epoch.Add(time.Hour - 1)))    //  true: 00:59:59.999...
	fmt.Println(span.Contains(epoch.Add(time.Hour)))        // false: 01:00:00.000...
	fmt.Println(span.Contains(epoch.Add(23*time.Hour - 1))) // false: 22:59:59.999...
	fmt.Println(span.Contains(epoch.Add(23 * time.Hour)))   //  true: 23:00:00.000...
	// Output:
	// true
	// false
	// false
	// true
}

func TestParse(t *testing.T) {
	save := time.Local
	time.Local, _ = time.LoadLocation("Asia/Tokyo")
	defer func() { time.Local = save }()

	tests := []struct {
		S string
		W string
	}{
		{"02:00+0000/2h", "02:00Z/2h0m0s"},
		{"12:34:56.123+0100/12h34m56s", "12:34:56.123+0100/12h34m56s"},

		{"02:00/2h", "02:00+0900/2h0m0s"},
		{"12:34:56.123/12h34m56s", "12:34:56.123+0900/12h34m56s"},

		{" 02:00/2h 14:00/2h ", "(02:00+0900/2h0m0s or 14:00+0900/2h0m0s)"},
	}
	for i, test := range tests {
		s, err := timespan.Parse(test.S)
		if err != nil {
			t.Error(i, test.S, err)
		}
		v := fmt.Sprint(s)
		if v != test.W {
			t.Error(i, test.S, v, test.W)
		}
	}

	t.Run("0h", func(t *testing.T) {
		span := timespan.MustParse("0h")
		epoch := time.Unix(0, 0)
		tests := []time.Time{
			epoch,
			epoch.Add(time.Hour - 1),
			epoch.Add(time.Hour),
			epoch.Add(23*time.Hour - 1),
			epoch.Add(23 * time.Hour),
		}
		for i, test := range tests {
			r := span.Contains(test)
			if r != false {
				t.Error(i, span, test)
			}
		}
	})

	t.Run("24h", func(t *testing.T) {
		span := timespan.MustParse("24h")
		epoch := time.Unix(0, 0)
		tests := []time.Time{
			epoch,
			epoch.Add(time.Hour - 1),
			epoch.Add(time.Hour),
			epoch.Add(23*time.Hour - 1),
			epoch.Add(23 * time.Hour),
		}
		for i, test := range tests {
			r := span.Contains(test)
			if r != true {
				t.Error(i, span, test)
			}
		}
	})
}
