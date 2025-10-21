package main

import "math"

// func Perimeter(width float64, height float64) float64 {
// 	return 2 * (width + height)
// }

// func Area(width float64, height float64) float64 {
// 	return width * height
// }

type Rectangle struct {
	Width  float64
	Height float64
}

// func (receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// func (receiverName ReceiverType) MethodName(args)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
