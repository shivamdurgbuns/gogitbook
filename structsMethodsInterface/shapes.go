package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	raidus float64
}

type Triangle struct {
	base   float64
	height float64
}

const (
	pi = math.Pi
)

type Shape interface {
	Area() float64
}

/*
Perimeter function returns the perimeter of a square/rectangle.
*/
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

/*
Rectangle Area method  gives the area of a square/rectangle
*/
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

/*
Circle Area method gives the area of a circle of radius r.
*/
func (c Circle) Area() float64 {
	return pi * c.raidus * c.raidus
}

/*
Triangel Area method gives the area of a triangle with base and height.
*/
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
