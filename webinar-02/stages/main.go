package main

import "fmt"

type LifeStage struct {
	YearFrom    uint
	YearTo      uint
	Title       string
	Description string
}

func (s LifeStage) Print() {
	fmt.Printf("%s (%d - %d)\n%s\n", s.Title, s.YearFrom, s.YearTo, s.Description)
}

func main() {
	s := LifeStage{
		YearFrom:    1999,
		YearTo:      2000,
		Title:       "Year of date time tragedy",
		Description: "A lot of code had broke due to 2000 year",
	}

	s.Print()
}
