package marsrover

type Commander interface {
	execute(rover IMarsRover)
}

type IMarsRover interface {
	moveForward()
	turnLeft()
	turnRight()
}