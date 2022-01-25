package mindiff3moves

import (
	"fmt"
	"sort"
)

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64

	minInt = -1 << (intSize - 1)
	maxInt = 1<<(intSize-1) - 1
)

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	// if len(nums) <= 8 {
	sort.Ints(nums)

	// firstHalf := nums[0 : len(nums)/2]
	// secondHalf := nums[len(nums)/2:]

	return minDifference3Moves(nums)

	// TODO: pick 4 largest and smallest

	// }

	// smallests := [4]int{maxInt, maxInt, maxInt, maxInt}
	// largests := [4]int{minInt, minInt, minInt, minInt}

	// for _, num := range nums {
	// 	for i, smallest := range smallests {
	// 		if num < smallest {
	// 			for j := i + 1; j < 3; j++ {
	// 				smallests[j] = smallests[j-1]
	// 			}

	// 			smallests[i] = num
	// 		}
	// 	}

	// 	for i, largest := range largests {
	// 		if num > largest {
	// 			for j := i + 1; j < 3; j++ {
	// 				largests[j] = largests[j-1]
	// 			}

	// 			largests[i] = num
	// 		}
	// 	}
	// }

	// return minDifference3Moves(smallests[:], largests[:])
}

func minDifference3Moves(all []int) int {
	// fmt.Printf("smallests: %+v\n", smallests)
	// fmt.Printf("largests: %+v\n", largests)

	minDiff := maxInt

	// all := append(smallests, largests...)
	// fmt.Printf("all: %+v\n", all)

	for i := range all {
		if i > 3 {
			break
		}

		for j := range all {
			if i+j > 3 {
				break
			}

			smallest := all[i]
			largest := all[len(all)-j-1]

			diff := largest - smallest
			if minDiff > diff {
				fmt.Printf("smallest: %v\n", smallest)
				fmt.Printf("largest: %v\n", largest)

				minDiff = diff
			}
		}
	}

	return minDiff
}
