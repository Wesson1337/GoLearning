package main

import "fmt"

func main() {
	stack := []int{1, 2, 3}
	stack = append(stack, 6)
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(top, stack)

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
