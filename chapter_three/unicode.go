package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 🔓 МИР"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 🔓 МИР" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)

	fmt.Println(string(65)) // "A", но не "65"
	fmt.Println(string(0x4eac)) // "C"

	fmt.Println(string(1234567))
}
