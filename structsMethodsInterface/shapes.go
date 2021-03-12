package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	raidus float64
}

const (
	pi = math.Pi
)

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
Circl Area method gives the area of a circle of radius r.
*/
func (c Circle) Area() float64 {
	return pi * c.raidus * c.raidus
}
