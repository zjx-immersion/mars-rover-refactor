package marsrover

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

var diractionMap = map[string]Direction{
	"N": N,
	"E": E,
	"S": S,
	"W": W,
}

func NewDirection(direction string) Direction {
	return Direction(valueOf(direction))
}

func valueOf(m string) Direction {
	return diractionMap[m]
}

func (t Direction) string() string {
	return [...]string{"N", "E", "S", "W"}[t]
}

func (m *Direction) turnRight() {
	//directionType := valueOf(m)
	newDirection := *m + 1
	if newDirection > 3 {
		newDirection = 0
	}
	*m = newDirection
}

func (m *Direction) turnLeft() {
	//directionType := valueOf(m)
	newDirection := *m - 1
	if newDirection < 0 {
		newDirection = 3
	}
	*m = newDirection
}
