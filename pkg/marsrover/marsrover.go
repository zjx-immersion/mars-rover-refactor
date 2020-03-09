package marsrover

type MarsRover struct {
	locationPointOfX int
	locationPointOfY int
	direction        string
}

var marsRover MarsRover

func initLocation(initLocationPointOfX, initLocationPointOfY int, initDirection string) {
	marsRover = MarsRover{initLocationPointOfX, initLocationPointOfY, initDirection}
}

func sendCommands(commands []byte) (int, int, string) {
	return marsRover.handleCommand(commands)
}

func (m *MarsRover) handleCommand(commands []byte) (int, int, string) {
	for _, cmd := range commands {
		if cmd == 'M' {
			m.moveForward()
		} else if cmd == 'L' {
			m.turnLeft()
		} else if cmd == 'R' {
			m.turnRight()
		}
	}

	return m.locationPointOfX, m.locationPointOfY, m.direction
}

func (m *MarsRover) turnRight() {
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

func (m *MarsRover) turnLeft() {
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

func (m *MarsRover) moveForward() {
	if m.direction == "E" {
		m.locationPointOfX += 1
	} else if m.direction == "W" {
		m.locationPointOfX -= 1
	} else if m.direction == "N" {
		m.locationPointOfY += 1
	} else if m.direction == "S" {
		m.locationPointOfY -= 1
	}
}
