package findwordsfromchars

func countCharacters(words []string, chars string) int {
	availableChars := map[rune]int{}

	for _, char := range chars {
		if _, found := availableChars[char]; !found {
			availableChars[char] = 0
		}

		availableChars[char]++
	}

	totalLen := 0

	for _, word := range words {
		requiredChars := map[rune]int{}

		for _, char := range word {
			if _, found := requiredChars[char]; !found {
				requiredChars[char] = 0
			}

			requiredChars[char]++
		}

		works := true

		for char, count := range requiredChars {
			if _, found := availableChars[char]; !found {
				works = false
				break
			}

			if count > availableChars[char] {
				works = false
				break
			}
		}

		if works {
			totalLen = totalLen + len(word)
		}
	}

	return totalLen
}
