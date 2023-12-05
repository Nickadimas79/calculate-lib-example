package triangle

import "math"

type Triangle struct {
	Height float32
	Base   float32
	SideA  float32
	SideB  float32
	SideC  float32
}

// Right function to find the area of a right triangle. Uses Base
// and Height for computation. Returns a float32.
//
// eq: (1/2) * (Base x Height)
func (t *Triangle) Right() float32 {
	return (1 / 2) * (t.Base * t.Height)
}

// Equilateral function to find the area of an equilateral triangle. Uses
// SideA as the side for computation. Returns a float32.
//
// eq: (√3)/4 x (SideA * SideA)
func (t *Triangle) Equilateral() float32 {
	rt := float32(math.Sqrt(3))

	return (rt / 4) * (t.SideA * t.SideA)
}

// Isosceles function to find the area of an isosceles triangle.
// Returns a float32.
//
// eq: (1/4) * Base * √(4 * (SideA * SideA) - (Base * Base))
func (t *Triangle) Isosceles() float32 {
	n := (1/4)*(t.SideA*t.SideA) - (t.Base * t.Base)
	rt := float32(math.Sqrt(float64(n)))

	return (1 / 4) * t.Base * rt
}

// SidesGiven function to find the area when all three sides are given.
// Returns a float32.
//
// eq:
func (t *Triangle) SidesGiven() float32 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	n := s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC)

	return float32(math.Sqrt(float64(n)))
}
