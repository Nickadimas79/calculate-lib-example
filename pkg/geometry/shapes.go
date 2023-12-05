package geometry

type Geometry interface {
	TwoD
	ThreeD
}

type TwoD interface {
	Area()
	Perimeter()
}

type ThreeD interface {
	Volume()
}
