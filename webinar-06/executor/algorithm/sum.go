package algorithm

type Sum struct{}

func (s Sum) GetName() string {
	return "sum"
}

func (s Sum) Calculate(nums []int) int {
	var sum int

	for _, n := range nums {
		sum += n
	}

	return sum
}
