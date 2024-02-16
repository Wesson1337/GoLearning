package main

func main() {
	x := "hello!"
	for _, x := range x {
		x := x + 'A' - 'a'
		println(string(x)) // "HELLO"
	}
}
