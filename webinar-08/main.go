package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex = &sync.Mutex{}
	var counter int

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Before lock, where i = ", i)

			mutex.Lock()

			fmt.Println("After lock, where i = ", i)

			counter += i

			mutex.Unlock()

			fmt.Println("After unlock, where i = ", i)

			fmt.Println(i, counter)
		}(i)
	}

	time.Sleep(time.Second)
}
