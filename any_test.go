package timespan

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	tests := []struct {
		span Span
		want bool
	}{
		{Or(Always), true},
		{Or(Never), false},
		{Or(Always, Always), true},
		{Or(Never, Always), true},
		{Or(Always, Never), true},
		{Or(Never, Never), false},
	}
	for i, test := range tests {
		r := test.span.Contains(time.Unix(0, 0))
		if r != test.want {
			t.Errorf("[%d] %v != %v", i, r, test.want)
		}
	}
}
