package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int

	go func() {
		fmt.Println("Hello!")
	}()

	for i := 0; i < 10; i++ {
		go func() {
			counter += i
			fmt.Println(i, counter)
		}()
	}

	time.Sleep(time.Second)
}
