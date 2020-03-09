package marsrover

type Direction struct {
	direction string
}

func NewDirection(direction string) Direction {
	return Direction{direction}
}

func (direction Direction) string() string {
	return direction.direction
}

func (m *Direction) turnRight() {
	if m.direction == "E" {
		m.direction = "S"
	} else if m.direction == "S" {
		m.direction = "W"
	} else if m.direction == "W" {
		m.direction = "N"
	} else if m.direction == "N" {
		m.direction = "E"
	}
}

func (m *Direction) turnLeft() {
	if m.direction == "E" {
		m.direction = "N"
	} else if m.direction == "N" {
		m.direction = "W"
	} else if m.direction == "W" {
		m.direction = "S"
	} else if m.direction == "S" {
		m.direction = "E"
	}
}
