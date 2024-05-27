package main

import "fmt"

type Duck interface {
	Quack() string
	Swim(in string) string
}

func tryDuck(d Duck) {
	fmt.Println("Quack:", d.Quack())
	fmt.Println("Swim in ocean:", d.Swim("ocean"))
}

type RareDuck struct{}

func (d RareDuck) Quack() string {
	return "Rare quack"
}

func (d RareDuck) Swim(in string) string {
	if in != "Royal lake" {
		return "I'm not swimming in here!"
	}

	return "Nice to swim here..."
}

type CommonDuck struct {
	Name string
}

func (d CommonDuck) Quack() string {
	return "I'm " + d.Name
}

func (d CommonDuck) Swim(in string) string {
	return "Cool! Swimming in " + in
}

func tryDucks() {
	tryDuck(RareDuck{})
	tryDuck(CommonDuck{"John"})
}
