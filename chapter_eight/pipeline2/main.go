package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sync()
	fmt.Println("sync - ", time.Since(now).Microseconds())

	now = time.Now()
	goroutine()
	fmt.Println("goroutine - ", time.Since(now).Microseconds())
}

func sync() {
	for x := 0; x < 100 ; x++ {
		natural := x
		square := natural * natural
		fmt.Println(square)
	}
}

func goroutine() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100 ; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- (x * x)
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}