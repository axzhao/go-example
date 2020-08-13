package basic

import (
	"testing"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

const (
	NORTH string = "North"
	EAST         = "East"
	SOUTH        = "South"
	WEST         = "West"
)

func (d Direction) String() string {
	return [...]string{NORTH, EAST, SOUTH, WEST}[d]
}

func TestEnum(t *testing.T) {
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
