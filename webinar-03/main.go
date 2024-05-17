package main

import "fmt"

type cup struct {
	size    string
	isClean bool
}

func (c *cup) clean() {
	c.isClean = true
	fmt.Println("Cleaned cup")
}

type teaCup struct {
	cup
	tea           string
	containsWater bool
}

func (c *teaCup) addTea(tea string) {
	c.tea = tea
	fmt.Printf("Added %s tea\n", tea)
}

func (c *teaCup) addWater() {
	c.containsWater = true
	fmt.Println("Added water")
}

func main() {
	createTeaCups()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Before continue", i)

	// 	if i < 3 {
	// 		continue
	// 	}

	// 	fmt.Println("After continue", i)
	// }
}

func createTeaCups() {
	var teaCounter = 0

	for teaCounter != 5 {
		// if teaCounter == 5 {
		// 	break
		// }

		c := teaCup{
			cup: cup{
				size: "BIG",
			},
		}

		c = prepareTea(c)

		teaCounter++
	}

	fmt.Println("FINISHED!")
}

func prepareTea(c teaCup) teaCup {
	if !c.isClean {
		c.clean()
	}

	tea := "black"

	c.addTea(tea)

	c.addWater()

	fmt.Printf("Got cup of tea: %+v\n", c)

	return c
}
