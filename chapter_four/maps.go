package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int) // Отображение строк на int
	fmt.Println(ages)

	ages = map[string]int {
		"alice": 31,
		"charlie": 34,
	}
	fmt.Println(ages)

	ages = make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)

	delete(ages, "alice")
	fmt.Println(ages)

	ages["bob"]++
	fmt.Println(ages)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var ages2 map[string]int
	fmt.Println(ages2 == nil) // "true"
	fmt.Println(len(ages2) == 0) // "true"

	age, ok := ages["bob"]
	if !ok {/*"bob" не является ключом в данном отображении; age ==0.*/}
	fmt.Println(age, ok)
}