package main

import (
	"fmt"
	"sync"
	"time"
)

func exampleDeadlock() {
	var ch chan int

	go func() {
		fmt.Println("Started go #1")
		i := <-ch
		fmt.Println(i)
		ch <- i
		fmt.Println("Finished go #1")
	}()

	go func() {
		fmt.Println("Started go #2")
		i := <-ch
		fmt.Println(i)
		fmt.Println("Finished go #2")
	}()

	time.Sleep(2 * time.Second)
}

func mutexExample() {
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
