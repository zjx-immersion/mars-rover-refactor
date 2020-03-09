package marsrover

type Coordinate struct {
	locationPointOfX int
	locationPointOfY int
}

func (c *Coordinate) move(coordinate Coordinate) {
	c.locationPointOfX += coordinate.locationPointOfX
	c.locationPointOfY += coordinate.locationPointOfY
}
