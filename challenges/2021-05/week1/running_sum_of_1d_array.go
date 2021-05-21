package week1

func runningSum(nums []int) []int {
	sums := make([]int, len(nums))

	for i, num := range nums {
		if i == 0 {
			sums[i] = num
			continue
		}

		sums[i] = num + sums[i-1]
	}

	return sums
}
