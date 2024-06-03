package batch

import (
	"fmt"
	"time"
	"webinar07/log"
)

type Inserter struct {
	source            log.Inserter
	ch                chan<- batch
	willCloseOnFinish chan struct{}
}

type batch struct {
	logs []log.Log
}

func NewInserter(s log.Inserter) *Inserter {
	ch := make(chan batch, 100)

	in := &Inserter{
		source:            s,
		ch:                ch,
		willCloseOnFinish: make(chan struct{}),
	}

	go runInsertion(in.source, ch, in.willCloseOnFinish)

	// go func() {}() // It also works!

	return in
}

func (in *Inserter) Insert(logs []log.Log) {
	fmt.Println("Sending batch via channel in Insert")

	in.ch <- batch{logs: logs}

	fmt.Println("Sent batch via channel in Insert")
}

func runInsertion(s log.Inserter, ch <-chan batch, closeOnFinish chan struct{}) {
	fmt.Println("Started runInsertion")

	var logsToInsert []log.Log

	ticker := time.NewTicker(2 * time.Second)

Loop:
	for {
		fmt.Println("Iterating in runInsertion")

		select {
		case <-ticker.C:
			fmt.Println("Got new tick")

			if len(logsToInsert) > 0 {
				s.Insert(logsToInsert)
				logsToInsert = nil
			}

		// send A to ch
		// send B to ch
		// send C to ch
		// close(ch)
		// cannot send anything to ch
		// read A, true    from ch
		// read B, true    from ch
		// read C, true    from ch
		// read nil, false from ch

		case b, ok := <-ch:
			if !ok {
				break Loop
			}

			fmt.Println("Got new value in runInsertion")

			logsToInsert = append(logsToInsert, b.logs...)

			if len(logsToInsert) >= 15 {
				s.Insert(logsToInsert)
				logsToInsert = nil
			}
		}
	}

	close(closeOnFinish) // TODO: check if it works
}

func (in *Inserter) Close() {
	close(in.ch)
	<-in.willCloseOnFinish
}
