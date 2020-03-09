package marsrover

import "mars-rover-refactor/pkg/marsrover/commands"

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

	var commanders = commands.ParseCommands(commandStrs)
	for _, commander := range commanders {
		commander.Execute(m)
	}
	return m.coordinate.locationPointOfX, m.coordinate.locationPointOfY, m.direction.string()
}

func (m *MarsRover) TurnRight() {
	m.direction.turnRight()
}

func (m *MarsRover) TurnLeft() {
	m.direction.turnLeft()
}

func (m *MarsRover) MoveForward() {
	newCoordinate := m.direction.produceForwardCoordinate()
	m.coordinate.move(newCoordinate)
}
