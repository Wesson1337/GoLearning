package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go spinner(500 * time.Millisecond)
	const n = 45
	fibN := fib(n) // медленно
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for i := range 3 {
			fmt.Print("\033[2K")
			fmt.Printf("\r%s", strings.Repeat(".", i+1))
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
