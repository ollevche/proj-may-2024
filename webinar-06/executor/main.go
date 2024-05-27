package main

import (
	"exec/algorithm"
	"exec/executor"
)

func main() {
	numbers := []int{1, 4, 29, 101, 4, 505}

	exec := executor.New().
		WithFormula(algorithm.Average{}).
		WithFormula(algorithm.Sum{}).
		WithFormula(algorithm.Lowest{})

	exec.Execute(numbers)
	exec.ExecuteByName("sum", numbers)
}
