package main

type Rectangle struct {
	width  float64
	height float64
}

/*
Perimeter function returns the perimeter of a square/rectangle.
*/
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

/*
Area functions gives the area of a square/rectangle
*/
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
