package main

import (
	"bytes"
	"strconv"
)

func main() {
	println(comma("-123.45")) // "12,345"

	s := "abc"
	b := []byte(s)
	s2 := string(b)
	println(s2)
}

// make this function non-recursive and use bytes.Buffer and working with float numbers

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	// Handle the sign separately
	sign := ""
	if s[0] == '+' || s[0] == '-' {
		sign = string(s[0])
		s = s[1:]
	}

	// Convert the string to float
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return ""
	}

	// Convert the float to string with comma
	s = strconv.FormatFloat(f, 'f', -1, 64)

	// Write the sign back to the buffer
	buf.WriteString(sign)

	for i, v := range s {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(v)
	}
	return buf.String()
}