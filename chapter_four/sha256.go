package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	a := []byte("x")
	b := []byte("X")
	var c1, c2 [32]byte
	var d1, d2 [64]byte

	if os.Args[1] == "256" {
		c1 = sha256.Sum256(a)
		c2 = sha256.Sum256(b)
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	} else {
		d1 = sha512.Sum512(a)
		d2 = sha512.Sum512(b)
		fmt.Printf("%x\n%x\n%t\n%T\n", d1, d2, d1 == d2, d1)
	}
	// fmt.Println(popCount(c1, c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

func popCount(sum1 [32]byte, sum2 [32]byte) int {
	var count int
	for i, v := range sum1 {
		if v != sum2[i] {
			count++
		} 	
	}
	return count
}