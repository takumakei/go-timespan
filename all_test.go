package timespan

import (
	"testing"
	"time"
)

func TestAnd(t *testing.T) {
	tests := []struct {
		span Span
		want bool
	}{
		{And(Always), true},
		{And(Never), false},
		{And(Always, Always), true},
		{And(Never, Always), false},
		{And(Always, Never), false},
		{And(Never, Never), false},
	}
	for i, test := range tests {
		r := test.span.Contains(time.Unix(0, 0))
		if r != test.want {
			t.Errorf("[%d] %v != %v", i, r, test.want)
		}
	}
}
