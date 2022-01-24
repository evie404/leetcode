package largestmagicsquare

func largestMagicSquare(grid [][]int) int {
	largestLen := 1

	shorterLen := len(grid)
	if len(grid[0]) < len(grid) {
		shorterLen = len(grid[0])
	}

	xSums := make([]int, shorterLen)
	ySums := make([]int, shorterLen)

	d1Sums := make([]int, shorterLen)

	for xStart := 0; xStart < len(grid); xStart++ {
		xLen := len(grid) - xStart

		for yStart := 0; yStart < len(grid[xStart]); yStart++ {
			yLen := len(grid[xStart]) - yStart

			maxLen := xLen
			if yLen < xLen {
				maxLen = yLen
			}

			// don't bother checking if there's already eq or bigger largestLen
			if maxLen <= largestLen {
				continue
			}

			candidateLens := make(map[int]struct{}, maxLen-largestLen)
			for testLen := largestLen + 1; testLen <= maxLen; testLen++ {
				candidateLens[testLen] = struct{}{}
			}

			xSums[0] = grid[xStart][yStart]
			ySums[0] = grid[xStart][yStart]
			d1Sums[0] = grid[xStart][yStart]

			for yOffset := 1; yOffset < maxLen; yOffset++ {
				xSums[yOffset] = xSums[yOffset-1] + grid[xStart][yStart+yOffset]
			}

			for xOffset := 1; xOffset < maxLen; xOffset++ {
				ySums[xOffset] = ySums[xOffset-1] + grid[xStart+xOffset][yStart]
			}

			for offset := 1; offset < maxLen; offset++ {
				d1Sums[offset] = d1Sums[offset-1] + grid[xStart+offset][yStart+offset]
			}

			// fmt.Printf("xStart: %v\n", xStart)
			// fmt.Printf("yStart: %v\n", yStart)
			// fmt.Printf("xSums: %+v\n", xSums)
			// fmt.Printf("ySums: %+v\n", ySums)
			// fmt.Printf("d1Sums: %+v\n", d1Sums)

			// fmt.Printf("candidateLens: %+v\n", candidateLens)

			for testLen := largestLen + 1; testLen <= maxLen; testLen++ {
				if _, isCandidate := candidateLens[testLen]; !isCandidate {
					continue
				}

				if xSums[testLen-1] != ySums[testLen-1] || xSums[testLen-1] != d1Sums[testLen-1] {
					delete(candidateLens, testLen)
					continue
				}

				// can probably cache this like d1Sum
				d2Sum := 0
				for offset := 0; offset < testLen; offset++ {
					d2Sum = d2Sum + grid[xStart+testLen-1-offset][yStart+offset]
				}

				if xSums[testLen-1] != d2Sum {
					delete(candidateLens, testLen)
					continue
				}

				rowCheck := true
				colCheck := true

				for xOffset := 0; xOffset < testLen; xOffset++ {
					currentSum := 0

					for yOffset := 0; yOffset < testLen; yOffset++ {
						currentSum = currentSum + grid[xOffset+xStart][yOffset+yStart]
					}

					if currentSum != xSums[testLen-1] {
						rowCheck = false
						break
					}
				}

				if !rowCheck {
					delete(candidateLens, testLen)
					continue
				}

				for yOffset := 0; yOffset < testLen; yOffset++ {
					currentSum := 0

					for xOffset := 0; xOffset < testLen; xOffset++ {
						currentSum = currentSum + grid[xOffset+xStart][yOffset+yStart]
					}

					if currentSum != xSums[testLen-1] {
						colCheck = false
						break
					}
				}

				if !colCheck {
					delete(candidateLens, testLen)
					continue
				}
			}

			// fmt.Printf("candidateLens: %+v\n", candidateLens)

			for candidate := range candidateLens {
				if candidate > largestLen {
					largestLen = candidate
				}
			}
		}
	}

	return largestLen
}
