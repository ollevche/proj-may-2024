package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) HappyBirthday() {
	// p.Age = p.Age + 1
	p.Age++
	fmt.Printf("%v is now %d years old\n", p.Name, p.Age)
}

// bad practice
func addTen(p *Person) {
	// p.Age = p.Age + 10
	p.Age += 10
}

func addTenSafe(p Person) Person {
	// p.Age = p.Age + 10
	p.Age += 10

	return p
}

func main() {
	p := Person{Name: "James"}

	p.HappyBirthday()

	fmt.Printf("%v, %v\n", p.Name, p.Age)

	addTen(&p)

	newP := addTenSafe(p)

	fmt.Printf("%v, %v\n", p.Name, p.Age)

	fmt.Printf("%v, %v\n", newP.Name, newP.Age)
}

func basicIntExample() {
	var i int = 2102

	// copy
	var c int = i

	i = 2103

	// no copy
	var p *int = &c

	fmt.Printf("i = %v\nc = %v\np = %v\n*p = %v\n", i, c, p, *p)

	*p = 101

	fmt.Println()

	fmt.Printf("i = %v\nc = %v\np = %v\n*p = %v\n", i, c, p, *p)
}
