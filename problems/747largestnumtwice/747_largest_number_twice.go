package largestnumbertwice

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64

	minInt = -1 << (intSize - 1)
)

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if len(nums) < 1 {
		return -1
	}

	largest := minInt
	largestIdx := 0
	secondLargest := minInt

	for i, num := range nums {
		if num > secondLargest && num < largest {
			secondLargest = num
		}

		if num > largest {
			secondLargest = largest
			largest = num
			largestIdx = i
		}
	}

	// fmt.Printf("largest: %v\n", largest)
	// fmt.Printf("secondLargest: %v\n", secondLargest)

	if largest >= 2*secondLargest {
		return largestIdx
	}

	return -1
}
