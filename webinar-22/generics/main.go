package main

import "fmt"

func main() {
	floats := []float64{1.1, 2.3, 5.1, 8.8, 10}
	ints := []int{1, 2, 3, -4, 5, -10}
	// unsignedInts := []uint{1, 4, 6, 90}
	strings := []string{"hello", "world"}

	printAllFloats(floats)
	printAllInts(ints)
	printAll(floats)
	printAll(ints)
	printAll(strings)

	fmt.Println(calculateSum(floats))
	fmt.Println(calculateSum(ints))
	// calculateSum(strings)
}

type Number interface {
	int | int64 | float64
}

func calculateSum[T Number](s []T) T {
	var sum T

	for _, se := range s {
		sum += se
	}

	return sum
}

func printAll[T any](s []T) {
	for _, se := range s {
		fmt.Println("Element is", se)
	}
}

func printAllFloats(s []float64) {
	for _, se := range s {
		fmt.Println("Element is", se)
	}
}

func printAllInts(s []int) {
	for _, se := range s {
		fmt.Println("Element is", se)
	}
}
