package main

import "fmt"


func main() {
	var runes []rune
	for _, r := range "Hello, мир" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'м' 'и' 'р']"

}