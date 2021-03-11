package main

/*
Sum function takes array of numbers and return the sum of the numbers in the array
*/
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
