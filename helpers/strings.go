package helpers

import (
	"fmt"
	"math"
	"strconv"
)

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

// Formats a duration specified in ticks.
// Assumes that the tick rate is 30 Hz.
func FormatTickToString(ticks int) string {
	const tickRate = 30

	float_seconds := float64(ticks) / tickRate
	centisecond := int(math.Floor(float_seconds*100)) % 100
	centisecond_str := strconv.Itoa(centisecond)
	if centisecond < 10 {
		centisecond_str = "0" + centisecond_str
	}

	seconds := (ticks % (tickRate * 60)) / tickRate
	seconds_str := strconv.Itoa(seconds)
	if seconds < 10 {
		seconds_str = "0" + seconds_str
	}

	minutes := ticks / (tickRate * 60)

	formated := fmt.Sprintf("%d:%s.%s", minutes, seconds_str, centisecond_str)
	return formated
}
