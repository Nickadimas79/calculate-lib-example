package circle

import "math"

type Circle struct {
	Radius float32
}

// Area returns the area of the circle object.
func (c *Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}
