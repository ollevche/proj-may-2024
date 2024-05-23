package main

import (
	"fmt"
	"hashtable/hashtable"
)

func main() {
	var h map[string]string

	fmt.Println(h["Marry"])

	h = make(map[string]string)

	h["John"] = "Software Engineer"

	h["James"] = "Plumber"

	h["Stacy"] = "Manager"

	fmt.Println(h["John"])
	fmt.Println(h["James"])
	fmt.Println(h["Stacy"])

	jv := h["John"]
	fmt.Println(jv)

	sv, ok := h["Stacy"]
	if ok {
		fmt.Println("Stacy is", sv)
	} else {
		fmt.Println("No info for Stacy")
	}

	delete(h, "Stacy")

	sv, ok = h["Stacy"]
	if ok {
		fmt.Println("Stacy is", sv)
	} else {
		fmt.Println("No info for Stacy")
	}

	fmt.Println("ITERATION")

	for k, v := range h {
		fmt.Println(k, v)
	}

	h["a"] = "b"

	for k, v := range h {
		fmt.Println(k, v)
	}
}

func exampleHashTable() {
	h := hashtable.New()

	h.Set("John", "Software Engineer")

	h.Set("James", "Plumber")

	h.Set("Stacy", "Manager")

	fmt.Println(h.Get("John"))
	fmt.Println(h.Get("James"))
	fmt.Println(h.Get("Stacy"))

	fmt.Println(h.Get("Mary"))

	fmt.Println("AFTER DELETE")

	h.Delete("Mary")

	h.Delete("John")

	if v, ok := h.Get("John"); ok {
		fmt.Println("John is ", v)
	} else {
		fmt.Println("No info for John")
	}
}
