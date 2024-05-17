package main

import (
	"fmt"
	"math/rand"
	"webinar3/check"
)

var globalI = 4

func pkgScope() {
	var wholeFunc = 2

	printGlobalI()
	if globalI == 4 {
		var ifonly = 101 + wholeFunc

		fmt.Println(ifonly)
	}
	// fmt.Println(ifonly)
	fmt.Println(check.IsBigger2(globalI))
	fmt.Println(check.GlobalS)
}

func printGlobalI() {
	fmt.Println(globalI)
}

func switchCase() {
	i := rand.Intn(100)

	fmt.Println(i)

	switch i {
	case 10:
		c := 1
		fmt.Println("10!", c)
	case 20:
		fmt.Println("20!")
	case 50:
		fmt.Println("50!")
	default:
		fmt.Println("not 10, 20, 50!")
	}

	switch {
	case i > 50:
		fmt.Println("> 50!")
	case i > 20:
		fmt.Println("> 20!")
	case i > 10:
		fmt.Println("> 10!")
	default:
		fmt.Println("< 10!")
	}

	switch j := i % 2; j {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("could not happen")
	}

	if j := i % 2; j == 0 {
		fmt.Println("j % 2 == 0")
	}

	// fmt.Println(j)
}

func ifelse() {
	i := rand.Intn(10)

	fmt.Println(i)

	if i > 5 {
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}

	if i >= 5 {
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}

	if i/2 == 2 {
		fmt.Println("/ 2 = 2")
	}

	if is4(i) && check.IsBigger2(i) {
		fmt.Println("4 and > 2")
	}

	// if (i == 2 || i == 4) && i < 10 || i == 5 {

	// }

	if is4(i) || (check.IsBigger2(i) && i < 1) {
		fmt.Println("4 or > 2")
	}
}

// func isBigger2(i int) bool {
// 	// if i > 2 {
// 	// 	return true
// 	// } else {
// 	// 	return false
// 	// }
// 	return i > 2
// }

func is4(i int) bool {
	if i == 4 {
		return true
	} else {
		return false
	}
}
