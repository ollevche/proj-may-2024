package batch

import (
	"context"
	"fmt"
	"sync"
	"time"
	"webinar07/log"
)

type Inserter struct {
	source            log.Inserter
	ch                chan batch
	willCloseOnFinish chan struct{}

	mutex             *sync.Mutex
	totalLogsInserted int
}

type batch struct {
	logs []log.Log
}

func NewInserter(ctx context.Context, s log.Inserter) *Inserter {
	ch := make(chan batch, 100)

	in := &Inserter{
		source:            s,
		ch:                ch,
		willCloseOnFinish: make(chan struct{}),
		mutex:             &sync.Mutex{},
	}

	go in.runInsertion(ctx, in.willCloseOnFinish)

	// go func() {}() // It also works!

	return in
}

func (in *Inserter) Insert(logs []log.Log) {
	fmt.Println("Sending batch via channel in Insert")

	in.ch <- batch{logs: logs}

	fmt.Println("Sent batch via channel in Insert")
}

func (in *Inserter) runInsertion(ctx context.Context, closeOnFinish chan struct{}) {
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
				in.insert(logsToInsert)
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

		case b, ok := <-in.ch:
			if !ok {
				break Loop
			}

			fmt.Println("Got new value in runInsertion")

			logsToInsert = append(logsToInsert, b.logs...)

			if len(logsToInsert) >= 15 {
				in.insert(logsToInsert)
				logsToInsert = nil
			}

		case <-ctx.Done():
			fmt.Println("CONTEXT is cancelled")
			in.insert(logsToInsert)
			break Loop
		}
	}

	close(closeOnFinish)
}

func (in *Inserter) insert(logs []log.Log) {
	in.source.Insert(logs)

	in.mutex.Lock()
	defer in.mutex.Unlock()

	in.totalLogsInserted += len(logs)
}

func (in *Inserter) GetTotalLogsInserted() int {
	in.mutex.Lock()
	defer in.mutex.Unlock()

	return in.totalLogsInserted
}

func (in *Inserter) ResetTotalLogsInserted() {
	in.mutex.Lock()
	defer in.mutex.Unlock()

	in.totalLogsInserted = 0
}

func (in *Inserter) Close() {
	close(in.ch)
	<-in.willCloseOnFinish
}
