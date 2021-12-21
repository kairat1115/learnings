package structsmethodsandinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectange struct {
	Width  float64
	Height float64
}

func (rectange Rectange) Perimeter() float64 {
	return 2 * (rectange.Width + rectange.Height)
}

func (rectange Rectange) Area() float64 {
	return rectange.Width * rectange.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) Area() float64 {
	return 0.5 * (triangle.Base * triangle.Height)
}
