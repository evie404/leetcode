package largestmagicsquare

func largestMagicSquare(grid [][]int) int {
	largest := 1
	// sums := make([][]int, len(grid))
	// for i := range sums {
	// 	sums[i] = make([]int, len(grid[i]))
	// }

	for xStart := 0; xStart < len(grid); xStart++ {
		xLen := len(grid) - xStart

		for yStart := 0; yStart < len(grid[xStart]); yStart++ {
			yLen := len(grid[xStart]) - yStart

			maxLen := xLen
			if yLen < xLen {
				maxLen = yLen
			}

			// don't bother checking if there's already eq or bigger largest
			if maxLen <= largest {
				continue
			}

			// candidateLens := make([]int, 0, maxLen)

			for testLen := largest + 1; testLen <= maxLen; testLen++ {
				subSquare := make([][]int, testLen)
				for i := range subSquare {
					subSquare[i] = make([]int, testLen)
				}

				for xOffset := 0; xOffset < testLen; xOffset++ {
					x := xStart + xOffset
					for yOffset := 0; yOffset < testLen; yOffset++ {
						y := yStart + yOffset

						subSquare[xOffset][yOffset] = grid[x][y]
					}
				}

				if isMagicSquare(subSquare) && testLen > largest {
					largest = testLen
				}
			}

			// for xOffset := 0; xOffset < maxLen; xOffset++ {
			// 	x := xStart + xOffset
			// 	for yOffset := 0; yOffset < maxLen; yOffset++ {
			// 		y := yStart + yOffset

			// 	}
			// }
		}
	}

	return largest
}

func isMagicSquare(grid [][]int) bool {
	if len(grid) < 1 {
		return false
	}

	if len(grid) != len(grid[0]) {
		return false
	}

	targetSum := 0

	for x := 0; x < len(grid); x++ {
		targetSum = targetSum + grid[x][0]
	}

	for x := 0; x < len(grid); x++ {
		currentSum := 0

		for y := 0; y < len(grid[x]); y++ {
			currentSum = currentSum + grid[x][y]
		}

		if currentSum != targetSum {
			return false
		}
	}

	for y := 0; y < len(grid[0]); y++ {
		currentSum := 0

		for x := 0; x < len(grid); x++ {
			currentSum = currentSum + grid[x][y]
		}

		if currentSum != targetSum {
			return false
		}
	}

	diagonalSum1 := 0
	diagonalSum2 := 0

	for x := 0; x < len(grid); x++ {
		diagonalSum1 = diagonalSum1 + grid[x][x]
		diagonalSum2 = diagonalSum2 + grid[len(grid)-x-1][x]
	}

	if diagonalSum1 != targetSum {
		return false
	}

	if diagonalSum2 != targetSum {
		return false
	}

	return true
}
