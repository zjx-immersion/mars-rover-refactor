package marsrover

//var locationPointOfX = 0
//var locationPointOfY = 0
//var direction = "N"

type MarsRover struct {
	locationPointOfX int
	locationPointOfY int
	direction        string
}

var marsRover MarsRover

func initLocation(initLocationPointOfX, initLocationPointOfY int, initDirection string) {
	marsRover = MarsRover{initLocationPointOfX, initLocationPointOfY, initDirection}
}

//func initLocation(initLocationPointOfX, initLocationPointOfY int, initDirection string) {
//	locationPointOfX = initLocationPointOfX
//	locationPointOfY = initLocationPointOfY
//	direction = initDirection
//}

func sendCommands(commands []byte) (int, int, string) {
	return marsRover.handleCommand(commands)
}

func (m *MarsRover) handleCommand(commands []byte) (int, int, string) {
	for _, cmd := range commands {
		if cmd == 'M' {
			if m.direction == "E" {
				m.locationPointOfX += 1
			} else if m.direction == "W" {
				m.locationPointOfX -= 1
			} else if m.direction == "N" {
				m.locationPointOfY += 1
			} else if m.direction == "S" {
				m.locationPointOfY -= 1
			}
		} else if cmd == 'L' {
			if m.direction == "E" {
				m.direction = "N"
			} else if m.direction == "N" {
				m.direction = "W"
			} else if m.direction == "W" {
				m.direction = "S"
			} else if m.direction == "S" {
				m.direction = "E"
			}
		} else if cmd == 'R' {
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
	}

	return m.locationPointOfX, m.locationPointOfY, m.direction
}

//func sendCommands(commands []byte) (int, int, string) {
//	for _, cmd := range commands {
//		if cmd == 'M' {
//			if direction == "E" {
//				locationPointOfX += 1
//			} else if direction == "W" {
//				locationPointOfX -= 1
//			} else if direction == "N" {
//				locationPointOfY += 1
//			} else if direction == "S" {
//				locationPointOfY -= 1
//			}
//		} else if cmd == 'L' {
//			if direction == "E" {
//				direction = "N"
//			} else if direction == "N" {
//				direction = "W"
//			} else if direction == "W" {
//				direction = "S"
//			} else if direction == "S" {
//				direction = "E"
//			}
//		} else if cmd == 'R' {
//			if direction == "E" {
//				direction = "S"
//			} else if direction == "S" {
//				direction = "W"
//			} else if direction == "W" {
//				direction = "N"
//			} else if direction == "N" {
//				direction = "E"
//			}
//		}
//	}
//
//	return locationPointOfX, locationPointOfY, direction
//}
