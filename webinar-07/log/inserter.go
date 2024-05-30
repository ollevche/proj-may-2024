package log

import "fmt"

type Inserter struct{}

// Insert emulates database insertion.
func (in Inserter) Insert(logs []Log) {
	fmt.Printf("<INSERTING %d LOGS>\n", len(logs))

	for _, l := range logs {
		fmt.Printf("ID = %v; V = %v\n", l.ID, l.Value)
	}

	fmt.Printf("<INSERTED %d LOGS>\n", len(logs))
}
