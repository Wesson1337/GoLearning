package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	x := 3
	now := time.Now()
	println(int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]))
	println(time.Since(now).Seconds())
	now = time.Now()
	sum := 0
	fmt.Printf("%X\n", pc)
	for i := range 7 {
		sum += int(pc[byte(x>>(i*8))])
	}
	println(sum)
	println(time.Since(now).Seconds())
}
