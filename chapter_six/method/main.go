package main

import (
	"fmt"
	"golearning/chapter_six/geometry"
)

type Point = geometry.Point

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(geometry.Distance(p, q)) // "5", вызов функции
	fmt.Println(p.Distance(q))           // "5",  вызов метода

	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance())

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p1 := Point{1, 2}
	pptr := &p1
	pptr.ScaleBy(2)
	fmt.Println(p1) // "{2, 4}"

	pptr.Distance(q) // implicit (*pptr).Distance(q)

	p2 := Point{1, 2}
	(&p2).ScaleBy(2)
	fmt.Println(p2) // "{2, 4}"

	p2.ScaleBy(2)   // implicit (&p2).ScaleBy(2)
	fmt.Println(p1) // "{4, 8}"

	// Point{1, 2}.ScaleBy(2) // Compile error: can't take address of Point literal

	il := IntList{1, &IntList{2, &IntList{3, nil}}}
	fmt.Println(il.Sum()) // "6"

	x := 5
	test(&x)
	fmt.Println(x) // "100"
}

// IntList is a linked list of integers.
// A nil *IntList represents an empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}


func test(lol *int) {
	*lol *= 20
}