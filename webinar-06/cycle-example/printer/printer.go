package printer

import (
	"cycle/paper"
	"fmt"
)

type Printer struct {
	p *paper.Paper
}

func (p *Printer) Store(paper paper.Paper) {
	p.p = &paper
}

func (p *Printer) Print() {
	if p.p != nil {
		fmt.Println(*p.p)
	} else {
		fmt.Println("NO PAPER")
	}
}
