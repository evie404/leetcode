package main

func mySqrt(x int) int {
	if x < 1 {
		return 0
	}

	i := 1

	for true {
		if i*i == x || i*i < x && (i+1)*(i+1) > x {
			return i
		}

		if i*i*4 < x {
			i = i * 2
		} else {
			i = i + 1
		}
	}

	return i
}
