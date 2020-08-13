package basic

import (
	"testing"
)

type Bits uint8

func Set(b, flag Bits) Bits       { return b | flag }
func Unset(b, flag Bits) Bits     { return b &^ flag }
func (b Bits) Has(flag Bits) bool { return b&flag != 0 }


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
