package training

type Rectangle struct {
	length float64
	width  float64
}

func (rect Rectangle) Perimeter() float64 {
	return (rect.length + rect.width) * 2
}

func (rect Rectangle) Surface() float64 {
	return rect.length * rect.width
}
