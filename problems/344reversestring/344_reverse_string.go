package main

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}

	var tmp byte

	for i := 0; i < len(s)/2; i++ {
		tmp = s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = tmp
	}
}
