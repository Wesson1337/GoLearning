package main

import (
	"fmt"
	"time"
)

func main() {
	bigSlowOperation()
	double(4) // "double(4) = 8"
	fmt.Println(triple(4)) // "12"
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	defer second()              // called second
	defer first()               // called first
	time.Sleep(5 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() { fmt.Printf("exit %s (%s)\n", msg, time.Since(start)) }
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
