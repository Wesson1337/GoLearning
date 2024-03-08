package main

import "fmt"


func main() {
	var runes []rune
	for _, r := range "Hello, мир" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'м' 'и' 'р']"

	var x []int
	x = append(x, 1)
	x = append(x, 1, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // Добавление среза x
	fmt.Println(x) // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}