package enum

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
