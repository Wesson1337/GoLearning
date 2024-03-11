package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(hypot(3, 4)) // "5"
	hypot(3, 4)

	fmt.Println(add(1, 2))   // "3"
	fmt.Println(sub(1, 2))   // "-1"
	fmt.Println(first(1, 2)) // "1"
	fmt.Println(zero(1, 2))  // "0"

	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"

	x := []int{1, 2, 3}
	fmt.Println(hello(x)) // "[1 2 3 1]"
	fmt.Println(x)        // "[1 2 3]"

	f := square
	fmt.Println(f(3)) // "9"

	f = negative
	fmt.Println(f(3))     // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"

	// f = product // compile error: can't assign func(int, int) int to func(int) int

	var f2 func(int) int
	// f2(3) // Аварийная остановка: вызов nil функции
	if f2 != nil {
		f(3)
	}

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}

func hypot(x, y float64) (z float64) {
	fmt.Println(z)
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func hello(x []int) []int {
	fmt.Println(x)
	j := append(x, 1)
	fmt.Println(x)
	return j
}

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }
