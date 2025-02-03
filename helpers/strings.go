package helpers

import (
	"bufio"
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
	buffer := bufio.NewReader(strings.NewReader(str))

	expandedChars := make([]ExpandedChar, 0)
	escapedChar := false
	var currentColor uint32
	for {
		char, _, err := buffer.ReadRune()
		if err != nil {
			break
		}

		if char == '#' && !escapedChar {
			colorBuffer := make([]byte, 8)
			_, err := buffer.Read(colorBuffer)
			if err != nil {
				break
			}

			color, err := strconv.ParseInt(string(colorBuffer), 16, 0)
			if err != nil {
				break
			}
			currentColor = uint32(color)
			continue
		}

		if char == '\\' && !escapedChar {
			escapedChar = true
			continue
		}

		escapedChar = false
		expandedChars = append(expandedChars, ExpandedChar{Char: char, Color: currentColor})
	}

	return expandedChars
}

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
