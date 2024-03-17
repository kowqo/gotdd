package main

import "math"

func main() {

}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

type Triangle struct {
	base, height float64
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}

type Shape interface {
	Area() float64
}
