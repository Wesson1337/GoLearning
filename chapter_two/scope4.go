package main

import "os"

func main() {
	if f, err := os.Open(fname); err != nil { // compile error: unused: f
		println(err)
	}
	f.Stat()  // compile error: undefined
	f.Close() // compile error: undefined

	// right:
	f, err := os.Open(fname)
	if err != nil {
		println(err)
	}
	f.Stat()
	f.Close()
}
