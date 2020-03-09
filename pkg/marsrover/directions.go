package marsrover

type Direction struct {
	directionDisplay string
}

type DirectionType int

const (
	N DirectionType = iota
	E
	S
	W
)

func NewDirection(direction string) Direction {
	return Direction{direction}
}

func valueOf(m *Direction) DirectionType {
	return map[string]DirectionType{
		"N": N,
		"E": E,
		"S": S,
		"W": W,
	}[m.directionDisplay]
}

func (t DirectionType) string() string {
	return [...]string{"N", "E", "S", "W"}[t]
}

func (direction Direction) string() string {
	return direction.directionDisplay
}

func (m *Direction) turnRight() {
	directionType := valueOf(m)
	newDirection := directionType + 1
	if newDirection > 3 {
		newDirection = 0
	}
	m.directionDisplay = newDirection.string()
}

func (m *Direction) turnLeft() {
	directionType := valueOf(m)
	newDirection := directionType - 1
	if newDirection < 0 {
		newDirection = 3
	}
	m.directionDisplay = newDirection.string()
}

func (m *Direction) turnRight2() {
	if m.directionDisplay == "E" {
		m.directionDisplay = "S"
	} else if m.directionDisplay == "S" {
		m.directionDisplay = "W"
	} else if m.directionDisplay == "W" {
		m.directionDisplay = "N"
	} else if m.directionDisplay == "N" {
		m.directionDisplay = "E"
	}
}

func (m *Direction) turnLeft2() {
	if m.directionDisplay == "E" {
		m.directionDisplay = "N"
	} else if m.directionDisplay == "N" {
		m.directionDisplay = "W"
	} else if m.directionDisplay == "W" {
		m.directionDisplay = "S"
	} else if m.directionDisplay == "S" {
		m.directionDisplay = "E"
	}
}
