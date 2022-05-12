package main

func majorityElement(nums []int) int {
	count := map[int]int{}

	max := 0
	maxCount := 0

	for _, num := range nums {
		if _, found := count[num]; !found {
			count[num] = 0
		}

		count[num]++

		if count[num] > maxCount {
			maxCount = count[num]
			max = num
		}
	}

	return max
}
