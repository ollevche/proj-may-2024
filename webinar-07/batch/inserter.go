package batch

import (
	"fmt"
	"time"
	"webinar07/log"
)

type Inserter struct {
	source log.Inserter
	ch     chan<- batch
}

type batch struct {
	logs []log.Log
}

func NewInserter(s log.Inserter) *Inserter {
	ch := make(chan batch, 1)

	in := &Inserter{
		source: s,
		ch:     ch,
	}

	go runInsertion(in.source, ch)

	// go func() {}() // It also works!

	return in
}

func (in Inserter) Insert(logs []log.Log) {
	fmt.Println("Sending batch via channel in Insert")

	in.ch <- batch{logs: logs}

	fmt.Println("Sent batch via channel in Insert")
}

func runInsertion(s log.Inserter, ch <-chan batch) {
	fmt.Println("Started runInsertion")

	var logsToInsert []log.Log

	ticker := time.NewTicker(2 * time.Second)

	for {
		fmt.Println("Iterating in runInsertion")

		select {
		case <-ticker.C:
			fmt.Println("Got new tick")

			if len(logsToInsert) > 0 {
				s.Insert(logsToInsert)
				logsToInsert = nil
			}

		case b := <-ch:
			fmt.Println("Got new value in runInsertion")

			logsToInsert = append(logsToInsert, b.logs...)

			if len(logsToInsert) >= 15 {
				s.Insert(logsToInsert)
				logsToInsert = nil
			}
		}
	}
}
