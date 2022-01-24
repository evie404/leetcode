package detectcapital

import "unicode"

func detectCapitalUse(word string) bool {
	if len(word) < 1 {
		return true
	}

	firstUpper := unicode.IsUpper(rune(word[0]))
	allUpper := unicode.IsUpper(rune(word[0]))
	allLower := unicode.IsLower(rune(word[0]))
	restLower := true

	for i := 1; i < len(word); i++ {
		character := rune(word[i])

		if unicode.IsUpper(character) {
			allLower = false
			restLower = false
		}

		if unicode.IsLower(character) {
			allUpper = false
		}
	}

	return allLower || allUpper || (firstUpper && restLower)
}
