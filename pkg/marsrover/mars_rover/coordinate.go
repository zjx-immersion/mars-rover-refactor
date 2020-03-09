package mars_rover

type Coordinate struct {
	locationPointOfX int
	locationPointOfY int
}


func NewCoordinate(locationPointOfX, locationPointOfY int) Coordinate {
	return Coordinate{locationPointOfX, locationPointOfY}
}

func (c *Coordinate) move(coordinate Coordinate) {
	c.locationPointOfX += coordinate.locationPointOfX
	c.locationPointOfY += coordinate.locationPointOfY
}
