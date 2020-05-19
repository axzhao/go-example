package bits

import (
	"testing"
)

func TestBits(t *testing.T) {
	const (
		first Bits = 1 << iota
		second
		third
		fourth
	)
	for _, r := range []struct {
		actual   Bits
		expected Bits
	}{
		{Set(first, second), Bits(3)},
		{Unset(Bits(3), second), Bits(1)},
	} {
		if r.actual != r.expected {
			t.Fatalf("Not equal: \n"+
				"expected: %v\n"+
				"actual: %v", r.expected, r.actual)
		}
	}
}
