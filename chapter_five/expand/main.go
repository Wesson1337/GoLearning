package main

import (
	"fmt"
	"strings"
)

func foo(s string) string {
	return s + "bar"
}

func main() {
	fmt.Println(expand("$foo", foo))
}

func expand(s string, f func(string) string) string {
	for {
		if strings.Contains(s, "$foo") {
			s = strings.Replace(s, "$foo", f("foo"), 1)
		} else {
			break
		}
	}
	return s
}
