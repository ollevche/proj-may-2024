package check

import "fmt"

var GlobalS = "Im global s"

func IsBigger2(i int) bool {
	return i > 2
}

func privateIsBigger2() {
	fmt.Println("im private privateIsBigger2")
}
