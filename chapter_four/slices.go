package main

import "fmt"

func main() {
	months := [...]string{1: "Январь", 12: "Декабрь"}
	fmt.Println(months)

	x := [3]int{1, 2, 3}
	c := x[:1]
	fmt.Println(c)
	fmt.Println(c[:])

	fmt.Printf("%q", months)
}
