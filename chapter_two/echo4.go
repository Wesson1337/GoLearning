package main

import (
	"flag"
	"fmt"
	"strings"
)

var global *int

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "разделитель")

func main() {
	var q int
	q = 1
	global = &q
	println(*global)
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println(*n)
	}
	t := 0
	println(t)
	test()
}

func test() {
	p := new(int)
	println("set")
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}