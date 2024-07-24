package helpers

// StripColorsFromString removes the colors of a given username.
func StripColorsFromString(name string) string {
	s := ""
	skipCount := 0
	for _, char := range name {
		// fmt.Printf("character %c starts at byte position %d\n", char, pos)
		if char == '#' {
			skipCount = 9
		}
		if skipCount == 0 {
			s = s + string(char)
		} else {
			skipCount--
		}
	}
	return s
}

// MostFrequentString finds the most frequent string in a string slice
func MostFrequentString(arr []string) string { // assuming no tie
	m := map[string]int{}
	var maxCnt int
	var mostFrequent string
	for _, a := range arr {
		m[a]++
		if m[a] > maxCnt {
			maxCnt = m[a]
			mostFrequent = a
		}
	}
	return mostFrequent
}

