package main

import (
	"fmt"
)

var b, f, s = true, 2.3, "four"
var x int = 292

func main() {
	k := &x
	p := *k
	fmt.Println(k)
	fmt.Println(p)

	var y int
	fmt.Println(&x == &x, &x == &y, &x == nil)

	a := 3
	fmt.Println(b, f, s, a)
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	// "32F = 0C"
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
	// "212F = 100C"

	var bar = foo()
	fmt.Println(bar, bar == foo())
	v := 1
	incr(&v)
	fmt.Println(v)
	fmt.Println(incr(&v))
}

func incr(count *int) int {
	*count++
	return *count
}

func foo() *int {
	v := 1
	return &v
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
