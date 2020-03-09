package marsrover

type MarsRover struct {
	coordinate Coordinate
	direction  Direction
}

func initLocation(initLocationPointOfX, initLocationPointOfY int, initDirection string) MarsRover {
	return MarsRover{
		Coordinate{initLocationPointOfX, initLocationPointOfY},
		NewDirection(initDirection)}
}

func (m *MarsRover) sendCommand(commandStrs []byte) (int, int, string) {

	var commanders = parseCommands(commandStrs)
	for _, commander := range commanders {
		commander.execute(m)
	}
	return m.coordinate.locationPointOfX, m.coordinate.locationPointOfY, m.direction.string()
}

func (m *MarsRover) turnRight() {
	m.direction.turnRight()
}

func (m *MarsRover) turnLeft() {
	m.direction.turnLeft()
}

func (m *MarsRover) moveForward() {
	newCoordinate := m.direction.produceForwardCoordinate()
	m.coordinate.move(newCoordinate)
}
