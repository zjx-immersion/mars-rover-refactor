package marsrover

type Direction string

func NewDirection(direction string) Direction {
	return Direction(direction)
}

func (direction Direction) string() string {
	return string(direction)
}

