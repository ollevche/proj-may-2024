package main

import (
	"fmt"
	"regexp"
)

func main() {
	// formattedString := "<p><b>Golang</b> <i>VS</i> <b>Python</b></p>"

	// pattern := regexp.MustCompile(`</?(b|i)>`)

	formattedString := "<i>VS</i>"

	pattern := regexp.MustCompile(`^<(p|b|i)>\w*</(p|b|i)>$`)

	results := pattern.FindAllString(formattedString, -1)
	for i, v := range results {
		fmt.Println(i, v)
	}
}

func basicExample() {
	matched, err := regexp.MatchString(`\w+`, "Cat meows")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(matched)

	pattern, err := regexp.Compile(`C\w+`)
	if err != nil {
		panic(err.Error())
	}

	matches := pattern.FindAllString("Cat named Cranchy meows", -1)
	for i, v := range matches {
		fmt.Println(i, v)
	}
}
