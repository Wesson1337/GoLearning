package main

import (
	"fmt"
)

func main() {
	ascii := 'a'
	unicode := '★'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "9733 ★ '★'"
	fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"
}
