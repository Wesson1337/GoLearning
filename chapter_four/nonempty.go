package main

import "fmt"

func main() {
	x := []string{"", "hello", "world"}
	fmt.Println(nonempty(x), x)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}