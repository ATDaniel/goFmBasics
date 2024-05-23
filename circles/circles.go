package circles

const (
	PI = 3.14
)

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) Circle {
	return Circle {
		Radius: radius,
	}
}

func (c *Circle) GetCircumference() float64 {
	return 2 * PI * c.Radius
}

func (c *Circle) GetArea() float64 {
	return PI * c.Radius * c.Radius
}