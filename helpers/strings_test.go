package helpers_test

import (
	// 	"bytes"
	// 	"common-go/helpers/strings"
	"reflect"
	"testing"

	common_helpers "github.com/pewpewlive/common-go/helpers"
)

func AssertEq(t *testing.T, expected any, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpected\n%v\nbut got\n%v", expected, actual)
	}
}

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

	AssertEq(t, expected, actual)
}

func TestStripColorsFromString(t *testing.T) {
	AssertEq(t, "", common_helpers.StripColorsFromString(""))
	AssertEq(t, "aaa", common_helpers.StripColorsFromString("aaa"))
	AssertEq(t, "", common_helpers.StripColorsFromString("#00112233"))
	AssertEq(t, "AAABBB", common_helpers.StripColorsFromString("AAA#00112233BBB"))
	AssertEq(t, "AAABBB", common_helpers.StripColorsFromString("AAA#00112233#44556677BBB"))
	AssertEq(t, "AAABBBCCC", common_helpers.StripColorsFromString("AAA#00112233BBB#44556677CCC"))
	// Incorrectly formatted strings
	AssertEq(t, "AAA", common_helpers.StripColorsFromString("AAA#0011"))
	AssertEq(t, "AAABBB", common_helpers.StripColorsFromString("AAA#0011#22334455BBB"))
}

func TestMostFrequentString(t *testing.T) {
	AssertEq(t, "foo", common_helpers.MostFrequentString([]string{"foo"}))
	AssertEq(t, "bar", common_helpers.MostFrequentString([]string{"foo", "bar", "bar"}))
	AssertEq(t, "bar", common_helpers.MostFrequentString([]string{"foo", "bar", "qux", "foo", "bar", "bar", "qux"}))
}

func TestFormatTickToString(t *testing.T) {
	// test centiseconds
	AssertEq(t, "0:00.00", common_helpers.FormatTicksToString(0))
	AssertEq(t, "0:00.03", common_helpers.FormatTicksToString(1))
	AssertEq(t, "0:00.06", common_helpers.FormatTicksToString(2))
	AssertEq(t, "0:00.10", common_helpers.FormatTicksToString(3))
	AssertEq(t, "0:00.50", common_helpers.FormatTicksToString(15))

	// test seconds
	AssertEq(t, "0:01.00", common_helpers.FormatTicksToString(30))
	AssertEq(t, "0:02.00", common_helpers.FormatTicksToString(30*2))
	AssertEq(t, "0:34.00", common_helpers.FormatTicksToString(30*34))

	// test minutes
	AssertEq(t, "1:00.00", common_helpers.FormatTicksToString(30*60))
	AssertEq(t, "4:00.00", common_helpers.FormatTicksToString(30*60*4))
	AssertEq(t, "56:00.00", common_helpers.FormatTicksToString(30*60*56))

	// everything together
	AssertEq(t, "2:34.56", common_helpers.FormatTicksToString(30*60*2+30*34+17))
}
