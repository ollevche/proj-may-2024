package algorithm

type Lowest struct{}

func (l Lowest) GetName() string {
	return "lowest"
}

func (l Lowest) Calculate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	v := nums[0]

	for _, n := range nums[1:] {
		if n < v {
			v = n
		}
	}

	return v
}
