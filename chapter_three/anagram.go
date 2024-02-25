package main

import "fmt"

func main() {
	fmt.Println(isAnagram("listen", "silent")) // true
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	mp := make(map[string]int)
	mp2 := make(map[string]int)
	for _, v := range s1 {
		mp[string(v)]++
	}
	for _, v := range s2 {
		mp2[string(v)]++
	}
	for k, v := range mp {
		if mp2[k] != v {
			return false
		}
	}
	return true
}