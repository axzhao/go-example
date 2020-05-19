package enum

import (
	"testing"
)

func TestBits(t *testing.T) {
	for _, r := range []struct {
		actual   Direction
		expected string
	}{
		{North, NORTH},
		{East, EAST},
		{South, SOUTH},
		{West, WEST},
	} {
		if r.actual.String() != r.expected {
			t.Fatalf("Not equal: \n"+
				"expected: %v\n"+
				"actual: %v", r.expected, r.actual)
		}
	}
}
