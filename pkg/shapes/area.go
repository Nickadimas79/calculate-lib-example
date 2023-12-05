package shapes

import "math"

func GetRectangleArea(width, length float32) float32 {
	return width * length
}

func GetRectangleVolume(length, width, height float32) float32 {
	return length * width * height
}

func GetCircleArea(radius float32) float32 {
	return math.Pi * radius * radius
}
