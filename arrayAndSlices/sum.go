package main

/*
Sum function takes array of numbers and return the sum of the numbers in the array
*/
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
