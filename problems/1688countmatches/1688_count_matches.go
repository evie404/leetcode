package countmatches

func numberOfMatches(teams int) int {
	matches := 0

	for teams > 1 {
		currentRoundMatches := teams / 2

		matches = matches + currentRoundMatches
		teams = teams - currentRoundMatches
	}

	return matches
}
