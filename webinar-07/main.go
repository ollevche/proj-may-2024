package main

import (
	"fmt"
	"math/rand/v2"
	"webinar07/log"
)

// Маємо певний тип LogInserter - він записує логи.
// Ціль - зробити LogInserter конкурентним та відправляти дані раз в N секунд.

func main() {
	var inserter log.Inserter

	for i := 0; i < 10; i++ {
		logs := make([]log.Log, 0, rand.IntN(10)+1)

		for j := 0; j < cap(logs); j++ {
			l := log.NewRandom(fmt.Sprintf("%d-%d", i, j))

			logs = append(logs, l)
		}

		inserter.Insert(logs)
	}
}
