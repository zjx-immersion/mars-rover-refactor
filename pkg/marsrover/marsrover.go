package marsrover

const forwardCommand = 'M'
const leftCommand = 'L'
const rightCommand = 'R'

type MarsRover struct {
	locationPointOfX int
	locationPointOfY int
	direction        string
}

type MarsRoverNew struct {
	coordinate Coordinate
	direction  string
}

func initLocation2(initLocationPointOfX, initLocationPointOfY int, initDirection string) MarsRover {
	return MarsRover{initLocationPointOfX, initLocationPointOfY, initDirection}
}

func initLocation(initLocationPointOfX, initLocationPointOfY int, initDirection string) MarsRoverNew {
	return MarsRoverNew{
		Coordinate{initLocationPointOfX, initLocationPointOfY},
		initDirection}
}

//func (m *MarsRover) sendCommand(commands []byte) (int, int, string) {
//	for _, cmd := range commands {
//		if cmd == byte(forwardCommand) {
//			m.moveForward()
//		} else if cmd == byte(leftCommand) {
//			m.turnLeft()
//		} else if cmd == byte(rightCommand) {
//			m.turnRight()
//		}
//	}
//	return m.locationPointOfX, m.locationPointOfY, m.direction
//}

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
	return m.coordinate.locationPointOfX, m.coordinate.locationPointOfY, m.direction
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

func (m *MarsRoverNew) moveForward() {
	if m.direction == "E" {
		m.coordinate.locationPointOfX += 1
	} else if m.direction == "W" {
		m.coordinate.locationPointOfX -= 1
	} else if m.direction == "N" {
		m.coordinate.locationPointOfY += 1
	} else if m.direction == "S" {
		m.coordinate.locationPointOfY -= 1
	}
}
