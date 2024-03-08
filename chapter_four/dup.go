package main

import "fmt"


func main() {
	x := []string{"hello", "h", "h", "w", "o", "w", "w", "w"}
	fmt.Println(dup(x))
}


// function that deletes all adjacent duplicates in a slice of strings
func dup(strings []string) []string {
	i := 0
	for _, s := range strings {
		if strings[i] != s {
			i++
			strings[i] = s
		}
	}
	return strings[:i+1]
}