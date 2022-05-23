package main

import "strconv"

func subtractProductAndSum(n int) int {
	s := strconv.Itoa(n)
	sum := 0
	mul := 1

	for _, c := range s {
		n := c - 48
		sum = sum + int(n)
		mul = mul * int(n)
	}

	return mul - sum
}
