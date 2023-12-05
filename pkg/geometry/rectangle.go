package geometry

type Rectangle struct {
	Length float32
	Width  float32
}

func (r *Rectangle) Area() float32 {
	return r.Width * r.Length
}
