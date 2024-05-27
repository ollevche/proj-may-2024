package executor

import (
	"fmt"
)

type Executor struct {
	formulas []Formula
}

type Formula interface {
	GetName() string
	Calculate([]int) int
}

func New() Executor {
	return Executor{}
}

func (e Executor) WithFormula(f Formula) Executor {
	e.formulas = append(e.formulas, f)
	return e
}

func (e Executor) Execute(numbers []int) {
	fmt.Println(numbers)

	for _, f := range e.formulas {
		result := f.Calculate(numbers)

		fmt.Printf("%s = %d\n", f.GetName(), result)
	}
}

func (e Executor) ExecuteByName(formulaName string, numbers []int) {
	fmt.Println(numbers)

	for _, f := range e.formulas {
		if f.GetName() != formulaName {
			continue
		}

		result := f.Calculate(numbers)

		fmt.Println(result)
	}
}
