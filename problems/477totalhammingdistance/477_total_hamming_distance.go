package totalhammingdistance

import (
	"fmt"
)

func totalHammingDistance(nums []int) int {
	total := 0

	// resultsCache := make(map[int]map[int]int, len(nums))

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			a := nums[i]
			b := nums[j]

			// smaller := a
			// larger := b
			// if b < a {
			// 	smaller = b
			// 	larger = a
			// }
			// var dist int

			// if cachedMap, foundMap := resultsCache[smaller]; foundMap {
			// 	if cachedDist, foundDist := cachedMap[larger]; foundDist {
			// 		dist = cachedDist
			// 	} else {
			// 		dist = hammingDistance(a, b)
			// 		resultsCache[smaller][larger] = dist
			// 	}
			// } else {
			// 	resultsCache[smaller] = make(map[int]int, len(resultsCache))
			// 	dist = hammingDistance(a, b)
			// 	resultsCache[smaller][larger] = dist
			// }

			dist := hammingDistance(a, b)

			total = total + dist
		}
	}

	return total
}

func hammingDistance(a, b int) int {
	if a == b {
		return 0
	}

	smaller := a
	larger := b
	if b < a {
		smaller = b
		larger = a
	}

	if smaller*2 == larger {
		return 2
	}

	aBin := fmt.Sprintf("%064b", a)
	bBin := fmt.Sprintf("%064b", b)

	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("b: %v\n", b)
	// fmt.Printf("aBin: %v\n", aBin)
	// fmt.Printf("bBin: %v\n", bBin)

	dist := 0

	for i := range aBin {
		if aBin[i] != bBin[i] {
			dist++
		}
	}

	return dist
}
