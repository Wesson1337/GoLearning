package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var arr []int
	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(1000000))
	}
	fmt.Println(len(arr))
	now := time.Now()
	sort.Ints(arr)
	fmt.Println(time.Since(now).Seconds())
}

// bubble sort
func bubble_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}