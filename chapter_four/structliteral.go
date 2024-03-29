package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	p := Point{1, 2}
	fmt.Println(p)

	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)                   // "false"

	p2 := Point{1, 2}
	fmt.Println(p == p2) // "true"
}
