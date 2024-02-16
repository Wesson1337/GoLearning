package main

import "fmt"

func main() {
	f := 3.141 // float64
	i := int(f)
	fmt.Println(f, i) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// Вывод:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
}
