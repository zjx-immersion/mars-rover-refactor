package commands

type Commander interface {
	Execute(rover IMarsRover)
}

type IMarsRover interface {
	MoveForward()
	TurnLeft()
	TurnRight()
}