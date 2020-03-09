package marsrover

type CommandParser struct{}

func parseCommands(commands []byte) []Commander {
	var commanders []Commander

	commanderMap := map[byte]Commander{
		'M': MoveForwardCommand{},
		'L': TurnLeftCommand{},
		'R': TurnRightCommand{},
	}

	for _, cmd := range commands {

		commanders = append(commanders, commanderMap[cmd])
	}
	return commanders
}

type TurnLeftCommand struct {
}

func (commander TurnLeftCommand) execute(rover IMarsRover) {
	rover.turnLeft()
}

type TurnRightCommand struct {
}

func (commander TurnRightCommand) execute(rover IMarsRover) {
	rover.turnRight()
}

type MoveForwardCommand struct {
}

func (commander MoveForwardCommand) execute(rover IMarsRover) {
	rover.moveForward()
}
