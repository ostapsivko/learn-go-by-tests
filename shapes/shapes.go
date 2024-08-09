package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return r.Width*2 + r.Height*2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rect Rectangle) float64 {
	return rect.Width*2 + rect.Height*2
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}
