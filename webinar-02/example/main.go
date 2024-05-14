package main

import "fmt"

type Age int

func main() {
	// name := Name("Oleksandr")
	// n := "Oleksandr"

	// age := 36

	person := Person{
		Name: "Oleksandr",
		Age:  36,
		// 20 more fields...
		Education: "KNEU",
	}

	person.print()
	person.printEducation()

	var nonamePerson Person

	nonamePerson.print()
	nonamePerson.printEducation()

	p := Person{"Stepan", 45, "PhD"}

	p.print()
	p.printEducation()

	// printPerson(name, Age(age))

	// name.print()
	// printName(name)
	// printName(Name(n))

	// fmt.Println("Hello world!")
}

type Name string

func (n Name) print() {
	fmt.Printf("The name is %v\n", n)
}

type Person struct {
	Name      Name
	Age       Age
	Education string
}

func (p Person) print() {
	p.Name.print()
	fmt.Printf("And their age is %v\n", p.Age)
}

func (p Person) printEducation() {
	fmt.Println(p.Education)
}
