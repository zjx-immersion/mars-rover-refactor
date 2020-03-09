package marsrover

var locationPointOfX = 0
var locationPointOfY = 0
var direction = "N"

func initLocation(initLocationPointOfX, initLocationPointOfY int, initDirection string) {
	locationPointOfX = initLocationPointOfX
	locationPointOfY = initLocationPointOfY
	direction = initDirection
}

func sendCommands(commands []byte) (int, int, string) {
	for _, cmd := range commands {
		if cmd == 'M' {
			if direction == "E" {
				locationPointOfX += 1
			} else if direction == "W" {
				locationPointOfX -= 1
			} else if direction == "N" {
				locationPointOfY += 1
			} else if direction == "S" {
				locationPointOfY -= 1
			}
		} else if cmd == 'L' {
			if direction == "E" {
				direction = "N"
			} else if direction == "N" {
				direction = "W"
			} else if direction == "W" {
				direction = "S"
			} else if direction == "S" {
				direction = "E"
			}
		} else if cmd == 'R' {
			if direction == "E" {
				direction = "S"
			} else if direction == "S" {
				direction = "W"
			} else if direction == "W" {
				direction = "N"
			} else if direction == "N" {
				direction = "E"
			}
		}
	}

	return locationPointOfX, locationPointOfY, direction
}
