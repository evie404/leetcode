package validmountainarray

func validMountainArray(arr []int) bool {
	downward := false
	upward := false

	if len(arr) < 3 {
		return false
	}

	for i, item := range arr {
		if i == 0 {
			continue
		}

		if item == arr[i-1] {
			return false
		}

		if item > arr[i-1] {
			if downward {
				return false
			}

			upward = true
			continue
		}

		downward = true
		continue
	}

	return downward && upward
}
