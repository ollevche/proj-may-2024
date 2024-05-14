package main

import "fmt"

func printHelloWorld(i int) {
	// var whatToPrint string

	whatToPrint := "Hello world!"

	// var counter int

	counter := 101 + 2

	// var half float64

	half := 0.5 + float64(counter)

	// var dollar rune

	dollar := '$'

	// var isMonday bool

	// isMonday = false

	println(whatToPrint, counter, half, dollar)

	fmt.Printf("%v %s %d", half, string(dollar), i)
}

func main() {
	fmt.Printf("%d", add5(3))
	// printHelloWorld(3)
}

func add5(number int) int {
	return number + 5
}
