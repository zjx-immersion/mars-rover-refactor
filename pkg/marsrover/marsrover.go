package marsrover

var x = 0
var y = 0
var direction = "N"

func initStatus(px, py int, pdirection string) {
	x = px
	y = py
	direction = pdirection
}

func run(commands []byte) (int, int, string) {
	for _, cmd := range commands {
		if cmd == 'M' {
			if direction == "E" {
				x += 1
			} else if direction == "W" {
				x -= 1
			} else if direction == "N" {
				y += 1
			} else if direction == "S" {
				y -= 1
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

	return x, y, direction
}
