package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	nonAnonymous()
	anonymous()
}

func nonAnonymous() {
	type Circle struct {
		Center Point
		Radius int
	}

	type Wheel struct {
		Circle Circle
		Spokes int
	}
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Println(w) // {{{8 8} 5} 20}
}

func anonymous() {
	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}
	var w Wheel
	w.X = 8      // w.Circle.Point.X = 8
	w.Y = 8      // w.Circle.Point.Y = 8
	w.Radius = 5 // w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Println(w) // {{{8 8} 5} 20}

	// literal
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // Примечание: завершающая запятая здесь (и после Radius) необходима
	}
	fmt.Printf("%#v\n", w)
	// Output:
	// main.Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// main.Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}
}
