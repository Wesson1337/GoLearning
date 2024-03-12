package main

import "fmt"

func main() {
	fmt.Println(sum()) // "0"
	fmt.Println(sum(3)) // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"

	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Printf("Line %d: ", linenum)
	fmt.Printf(format, args...)
	fmt.Println()
}

func f(...int) {} 
func g([]int) {}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	min := vals[0]
	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}
	return min
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals[1:] {
		if val > max {
			max = val
		}
	}
	return max
}

func strjoin(sep string, strs ...string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}