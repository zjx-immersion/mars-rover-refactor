package marsrover

type Direction int

const (
	N Direction = iota
	E
	S
	W
)

var directionDisplays = [...]string{"N", "E", "S", "W"}

var diractionMap = map[string]Direction{
	directionDisplays[N]: N,
	directionDisplays[E]: E,
	directionDisplays[S]: S,
	directionDisplays[W]: W,
}

func NewDirection(direction string) Direction {
	return Direction(valueOf(direction))
}

func valueOf(m string) Direction {
	return diractionMap[m]
}

func (d Direction) string() string {

	return directionDisplays[d]
}

func (d *Direction) turnRight() {
	//directionType := valueOf(d)
	newDirection := *d + 1
	if newDirection > 3 {
		newDirection = 0
	}
	*d = newDirection
}

func (d *Direction) turnLeft() {
	//directionType := valueOf(d)
	newDirection := *d - 1
	if newDirection < 0 {
		newDirection = 3
	}
	*d = newDirection
}

func (d *Direction) produceForwardCoordinate() Coordinate {
	directionMoveCoordinateMap := map[Direction]Coordinate{
		E: NewCoordinate(1, 0),
		W: NewCoordinate(-1, 0),
		N: NewCoordinate(0, 1),
		S: NewCoordinate(0, -1),
	}
	return directionMoveCoordinateMap[*d]
}
