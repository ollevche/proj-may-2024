package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
	"webinar07/batch"
	"webinar07/log"
)

// Маємо певний тип LogInserter - він записує логи.
// Ціль - зробити LogInserter конкурентним та відправляти дані раз в N секунд.

func main() {
	var inserter log.Inserter

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	batchInserter := batch.NewInserter(ctx, inserter)

	go func() {
		t := time.NewTicker(time.Second / 2)
		lastResetAt := time.Now()

		for {
			fmt.Println("Before <-t.C")

			<-t.C

			fmt.Println("Total inserted: ", batchInserter.GetTotalLogsInserted())

			if time.Since(lastResetAt) >= time.Second {
				fmt.Println("Reseting...")
				batchInserter.ResetTotalLogsInserted()
			}
		}
	}()

	generateAndInsertLogs(batchInserter)

	time.Sleep(time.Second * 2)

	batchInserter.Close()
}

type Inserter interface {
	Insert(logs []log.Log)
}

func generateAndInsertLogs(inserter Inserter) {
	for i := 0; i < 10; i++ {
		logs := make([]log.Log, 0, rand.IntN(10)+1)

		for j := 0; j < cap(logs); j++ {
			l := log.NewRandom(fmt.Sprintf("%d-%d", i, j))

			logs = append(logs, l)
		}

		inserter.Insert(logs)
	}
}

func exampleGoTo() {
LABEL_1:
	fmt.Println("Hello")

	goto LABEL_1
}
