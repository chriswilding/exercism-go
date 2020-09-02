package romannumerals

import (
	"errors"
)

var arabicToRoman = []struct {
	arabic int
	roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("number must be between 1 and 3000 inclusive")
	}

	var roman string

	for arabic > 0 {
		for _, atr := range arabicToRoman {
			if arabic >= atr.arabic {
				arabic -= atr.arabic
				roman += atr.roman
				break
			}
		}
	}

	return roman, nil
}
