package main

import "fmt"

func main() {
	const (
		e  = 2.718288
		pi = 3.14159 // Лучшим приближением является math.Pi
	)

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"


	type Weekday int
	const (
		Sunday Weekday = iota // 0
		Monday // 1
		Tuesday // 2
		Wednesday // 3
		Thursday // 4
		Friday // 5
		Saturday // 6
	)

	type Flags uint
	const (
		FlagUp = 1 << iota // 1 = 1
		FlagBroadcast // 2 = 10
		FlagLoopBack // 4 = 100
		FlagPointToPoint // 8 = 1000
		FlagMulticast // 16 = 10000
	)

	const (
		_ = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
	)

	var f float64 = 3 + 0i
	println(f)
}