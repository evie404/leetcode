package repeatedstringmatch

import (
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	parts := strings.Split(b, a)

	if strings.Contains(a, b) {
		return 1
	}

	// fmt.Printf("parts: %+v\n", parts)
	// fmt.Printf("len(parts): %v\n", len(parts))
	// fmt.Printf("middle: %+v\n", middle)

	if len(parts) < 1 {
		return -1
	}

	if len(parts) == 1 {
		if strings.HasSuffix(a, parts[0]) || strings.HasPrefix(a, parts[0]) {
			return 1
		}

		if strings.Contains(a+a, parts[0]) {
			return 2
		}
	}

	middleAllEmpty := true

	middle := parts
	if len(parts) >= 2 {
		middle = parts[1 : len(parts)-1]
	}

	for _, part := range middle {
		if part != "" {
			middleAllEmpty = false
			break
		}
	}

	if !middleAllEmpty {
		return -1
	}

	if parts[0] == "" && parts[len(parts)-1] == "" {
		return len(parts) - 1
	}

	if parts[0] == "" && strings.HasPrefix(a, parts[len(parts)-1]) {
		return len(parts)
	}

	if parts[len(parts)-1] == "" && strings.HasSuffix(a, parts[0]) {
		return len(parts)
	}

	if strings.HasSuffix(a, parts[0]) && strings.HasPrefix(a, parts[len(parts)-1]) {
		return len(parts) + 1
	}

	return -1
}
