package word

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		for i < j && !isChar(runes[i]) {
			i++
		}
		for i < j && !isChar(runes[j]) {
			j--
		}
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isChar(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}