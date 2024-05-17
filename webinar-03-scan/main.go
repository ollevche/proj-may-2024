package main

import "fmt"

func main() {
	fmt.Println("Enter your year of birth:")

	var year int

	// Scan function reads user input from terminal and stores
	// it into year variable by pointer.
	// Docs: https://pkg.go.dev/fmt#Scan
	fmt.Scan(&year)

	if year > 2024 {
		fmt.Println("Hello, time traveler!")
	} else {
		fmt.Printf("Seems like you are %d years old!\n", 2024-year)
	}
}
