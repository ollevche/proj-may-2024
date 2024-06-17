package main

import "fmt"

func main() {
	runDivideByZero()
	fmt.Println("4")
}

func runDivideByZero() {
	defer func() {
		v := recover()
		fmt.Println(v)
	}()

	divideByZero()

	fmt.Println("3")
}

func divideByZero() {
	fmt.Println("1")

	panic("Procrastinate")

	fmt.Println("2")
}
