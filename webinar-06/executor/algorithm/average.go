package algorithm

type Average struct{}

func (a Average) GetName() string {
	return "average"
}

func (a Average) Calculate(nums []int) int {
	var sum int

	for _, n := range nums {
		sum += n
	}

	average := sum / len(nums)

	return average
}
