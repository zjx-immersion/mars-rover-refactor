package marsrover

const forwardCommand = 'M'
const leftCommand = 'L'
const rightCommand = 'R'

type MarsRover struct {
	coordinate Coordinate
	direction  Direction
}

func initLocation(initLocationPointOfX, initLocationPointOfY int, initDirection string) MarsRover {
	return MarsRover{
		Coordinate{initLocationPointOfX, initLocationPointOfY},
		NewDirection(initDirection)}
}

func (m *MarsRover) sendCommand(commands []byte) (int, int, string) {
	for _, cmd := range commands {
		if cmd == byte(forwardCommand) {
			m.moveForward()
		} else if cmd == byte(leftCommand) {
			m.direction.turnLeft()
		} else if cmd == byte(rightCommand) {
			m.direction.turnRight()
		}
	}
	return m.coordinate.locationPointOfX, m.coordinate.locationPointOfY, m.direction.string()
}

func (m *MarsRover) moveForward() {
	newCoordinate := m.direction.produceForwardCoordinate()
	m.coordinate.move(newCoordinate)
}

