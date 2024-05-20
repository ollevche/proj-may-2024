package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	Name    string
	IsAdmin bool
}

func (u User) print() {
	fmt.Println(u.Name, u.IsAdmin)

	if u.IsAdmin {
		fmt.Println("User is admin")
	}
}

func main() {
	var users []User

	for i := 0; i < 10; i++ {
		users = append(users, User{
			Name:    fmt.Sprintf("User #%d", i),
			IsAdmin: rand.Intn(2) == 1,
		})
	}

	for _, u := range users { // for <index>, <value> := range <data_structure>
		u.print()
	}
}

func deleteByIndex(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)
}

func exampleDelete() {
	abc := []string{"a", "b", "c", "d", "e"}

	fmt.Println(deleteByIndex(abc, 3))
}

func exampleSlicing() {
	var firstSlice = make([]string, 4)

	for i := range firstSlice {
		firstSlice[i] = "banana"
	}

	fmt.Println("1. firstSlice", firstSlice)

	var secondSlice = firstSlice[:] // 0:len(firstSlice)

	secondSlice[1] = "apple"

	fmt.Println("2. firstSlice", firstSlice)
	fmt.Println("2. secondSlice", secondSlice)

	secondSlice = secondSlice[1:]
	fmt.Println("3. firstSlice", firstSlice)
	fmt.Println("3. secondSlice", secondSlice)

	// a[low:high:max]
	// The indices low and high select which elements of operand a appear in the result.
	// Max controls the resulting slice's capacity by setting it to max - low.
	// Language specs: https://go.dev/ref/spec#Slice_expressions (Full slice expressions)
	thirdSlice := firstSlice[2:4:4]
	fmt.Println("4. thirdSlice", thirdSlice, cap(thirdSlice))
}

func exampleSlice() {
	// CamelCase -> public
	// camelCase -> private

	var sliceLiteral []int

	sliceLiteral = []int{101, 202, 303}

	fmt.Println(sliceLiteral, len(sliceLiteral), cap(sliceLiteral))

	sliceMake := make([]int, 1, 3) // make(type, len, cap) OR make(type, len)

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	sliceMake[0] = 101

	pointerToFirstElem := &sliceMake[0]

	fmt.Println("BY POINTER:", *pointerToFirstElem)

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	sliceMake = append(sliceMake, 202, 303, 404)

	*pointerToFirstElem = 55

	if sliceMake[0] == 55 {
		sliceMake[0] = 505
	}

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))
}

func exampleArray() {
	var arr [3]int

	arr[0] = 1

	arr[0] = 101

	setArrToFives(arr)

	printArr(arr)

	fmt.Println(arr)

	fmt.Println(len(arr), cap(arr))
}

func setArrToFives(arr [3]int) {
	arr[0] = 5
	arr[1] = 5
	arr[2] = 5
}

func printArr(arr [3]int) {
	for i, v := range arr {
		v = 202
		arr[i] = 202
		fmt.Println(i, arr[i], v)
	}
}
