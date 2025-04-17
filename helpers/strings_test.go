package helpers_test

import (
	// 	"bytes"
	// 	"common-go/helpers/strings"

	"testing"

	common_helpers "github.com/pewpewlive/common-go/helpers"
	"github.com/stretchr/testify/assert"
)

func TestStringToExpandedChars(t *testing.T) {
	expected := []common_helpers.ExpandedChar{
		{Char: 'M', Color: 0xffff00ff},
		{Char: 'ū', Color: 0xffff00ff},
		{Char: 's', Color: 0xffff00ff},
		{Char: 'ų', Color: 0xffff00ff},
		{Char: 'Š', Color: 0x007fffff},
		{Char: 'o', Color: 0x007fffff},
		{Char: 'k', Color: 0xffff00ff},
		{Char: 'i', Color: 0x00ffffff},
		{Char: 'ų', Color: 0x000000ff},
		{Char: 'P', Color: 0xffff00ff},
		{Char: 'a', Color: 0xffff00ff},
		{Char: 'm', Color: 0xffff00ff},
	}

	actual := common_helpers.StringToExpandedChars("#ffff00ffMūsų#007fffffŠo#ffff00ffk#00ffffffi#000000ffų#ffff00ffPam")

	assert.Equal(t, expected, actual)
}

func TestStripColorsFromString(t *testing.T) {
	assert.Equal(t, "", common_helpers.StripColorsFromString(""))
	assert.Equal(t, "aaa", common_helpers.StripColorsFromString("aaa"))
	assert.Equal(t, "", common_helpers.StripColorsFromString("#00112233"))
	assert.Equal(t, "AAABBB", common_helpers.StripColorsFromString("AAA#00112233BBB"))
	assert.Equal(t, "AAABBB", common_helpers.StripColorsFromString("AAA#00112233#44556677BBB"))
	assert.Equal(t, "AAABBBCCC", common_helpers.StripColorsFromString("AAA#00112233BBB#44556677CCC"))
	// Incorrectly formatted strings
	assert.Equal(t, "AAA", common_helpers.StripColorsFromString("AAA#0011"))
	assert.Equal(t, "AAABBB", common_helpers.StripColorsFromString("AAA#0011#22334455BBB"))
}

func TestMostFrequentString(t *testing.T) {
	assert.Equal(t, "foo", common_helpers.MostFrequentString([]string{"foo"}))
	assert.Equal(t, "bar", common_helpers.MostFrequentString([]string{"foo", "bar", "bar"}))
	assert.Equal(t, "bar", common_helpers.MostFrequentString([]string{"foo", "bar", "qux", "foo", "bar", "bar", "qux"}))
}

func TestFormatTickToString(t *testing.T) {
	// test centiseconds
	assert.Equal(t, "0:00.00", common_helpers.FormatTicksToString(0))
	assert.Equal(t, "0:00.03", common_helpers.FormatTicksToString(1))
	assert.Equal(t, "0:00.06", common_helpers.FormatTicksToString(2))
	assert.Equal(t, "0:00.10", common_helpers.FormatTicksToString(3))
	assert.Equal(t, "0:00.50", common_helpers.FormatTicksToString(15))

	// test seconds
	assert.Equal(t, "0:01.00", common_helpers.FormatTicksToString(30))
	assert.Equal(t, "0:02.00", common_helpers.FormatTicksToString(30*2))
	assert.Equal(t, "0:34.00", common_helpers.FormatTicksToString(30*34))

	// test minutes
	assert.Equal(t, "1:00.00", common_helpers.FormatTicksToString(30*60))
	assert.Equal(t, "4:00.00", common_helpers.FormatTicksToString(30*60*4))
	assert.Equal(t, "56:00.00", common_helpers.FormatTicksToString(30*60*56))

	// everything together
	assert.Equal(t, "2:34.56", common_helpers.FormatTicksToString(30*60*2+30*34+17))
}
