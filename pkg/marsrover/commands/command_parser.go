package commands

type CommandParser struct{}

var commanderMap = map[byte]Commander{
	'M': MoveForwardCommand{},
	'L': TurnLeftCommand{},
	'R': TurnRightCommand{},
}

func ParseCommands(commands []byte) []Commander {
	var commanders []Commander

	for _, cmd := range commands {

		commanders = append(commanders, commanderMap[cmd])
	}
	return commanders
}

type TurnLeftCommand struct {
}

func (commander TurnLeftCommand) Execute(rover IMarsRover) {
	rover.TurnLeft()
}

type TurnRightCommand struct {
}

func (commander TurnRightCommand) Execute(rover IMarsRover) {
	rover.TurnRight()
}

type MoveForwardCommand struct {
}

func (commander MoveForwardCommand) Execute(rover IMarsRover) {
	rover.MoveForward()
}
