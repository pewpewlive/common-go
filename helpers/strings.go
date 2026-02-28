package helpers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ExpandedChar struct {
	Char  rune
	Color uint32
}

func StringToExpandedChars(str string) []ExpandedChar {
	expandedChars := make([]ExpandedChar, 0)
	var currentColor uint32 = 0xffffffff
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		if char == '#' {
			if i+8 < len(runes) {
				colorStr := string(runes[i+1 : i+9])
				color, err := strconv.ParseUint(colorStr, 16, 32)
				if err == nil {
					currentColor = uint32(color)
				} else {
					currentColor = 0xffffffff
				}
				i += 8
				continue
			} else {
				break
			}
		}
		expandedChars = append(expandedChars, ExpandedChar{Char: char, Color: currentColor})
	}
	return expandedChars
}

func ExpandedCharsToString(expandedChars []ExpandedChar) string {
	var builder strings.Builder
	var currentColor uint32 = 0xffffffff
	for _, expandedChar := range expandedChars {
		if expandedChar.Color != currentColor {
			builder.WriteString(fmt.Sprintf("#%08x", expandedChar.Color))
			currentColor = expandedChar.Color
		}
		builder.WriteRune(expandedChar.Char)
	}
	return builder.String()
}

// StripColorsFromString removes the colors of a given username.
func StripColorsFromString(name string) string {
	expandedChars := StringToExpandedChars(name)
	var builder strings.Builder
	for _, expandedChar := range expandedChars {
		builder.WriteRune(expandedChar.Char)
	}
	return builder.String()
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
func FormatTicksToString(ticks int64) string {
	const tickRate = 30

	float_seconds := float64(ticks) / tickRate
	centisecond := int(math.Floor(float_seconds*100)) % 100
	centisecond_str := strconv.Itoa(centisecond)
	if centisecond < 10 {
		centisecond_str = "0" + centisecond_str
	}

	seconds := (ticks % (tickRate * 60)) / tickRate
	seconds_str := strconv.FormatInt(seconds, 10)
	if seconds < 10 {
		seconds_str = "0" + seconds_str
	}

	minutes := ticks / (tickRate * 60)

	formated := fmt.Sprintf("%d:%s.%s", minutes, seconds_str, centisecond_str)
	return formated
}
