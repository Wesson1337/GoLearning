package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	q = [3]int{1, 2, 3}
	// q = [4]int{1, 2, 3, 4} // Ошибка компиляции: нельзя присвоить [4]int переменной типа [3]int

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RUR
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RUR: "₽"}
	fmt.Println(RUR, symbol[RUR]) // "₽"

	j := [...]int{99:-1}
	fmt.Println(j[0], j[99]) // "0" "-1"

	x := [2]int{1, 2}
	y := [...]int{1, 2}
	z := [2]int{1, 3}
	fmt.Println(x == y, x == z, y == z) // "true false false"
	// d := [3]int{1, 2}
	// fmt.Println(x == d) // Ошибка компиляции: разные типы [2]int и [3]int

	m := [32]byte{31:1}
	fmt.Println(m[31]) // "1"
	zero(&m)
	fmt.Println(m[31]) // "0"
}


func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
