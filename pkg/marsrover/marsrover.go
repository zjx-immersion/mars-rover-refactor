package marsrover

const forwardCommand = 'M'
const leftCommand = 'L'
const rightCommand = 'R'

type MarsRoverNew struct {
	coordinate Coordinate
	direction  Direction
}

func initLocation(initLocationPointOfX, initLocationPointOfY int, initDirection string) MarsRoverNew {
	return MarsRoverNew{
		Coordinate{initLocationPointOfX, initLocationPointOfY},
		NewDirection(initDirection)}
}

func (m *MarsRoverNew) sendCommand(commands []byte) (int, int, string) {
	for _, cmd := range commands {
		if cmd == byte(forwardCommand) {
			m.moveForward()
		} else if cmd == byte(leftCommand) {
			m.turnLeft()
		} else if cmd == byte(rightCommand) {
			m.turnRight()
		}
	}
	return m.coordinate.locationPointOfX, m.coordinate.locationPointOfY, m.direction.string()
}

func (m *MarsRoverNew) turnRight() {
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

func (m *MarsRoverNew) turnLeft() {
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

func (m *MarsRoverNew) moveForward() {
	directionMoveCoordinateMap := map[string]Coordinate{
		"E": NewCoordinate(1, 0),
		"W": NewCoordinate(-1, 0),
		"N": NewCoordinate(0, 1),
		"S": NewCoordinate(0, -1),
	}
	m.coordinate.move(directionMoveCoordinateMap[m.direction.string()])
}

func NewCoordinate(locationPointOfX, locationPointOfY int) Coordinate {
	return Coordinate{locationPointOfX, locationPointOfY}
}
