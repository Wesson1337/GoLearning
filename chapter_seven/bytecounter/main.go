package main

import (
	"fmt"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	words := strings.Fields(string(p))
	*c += WordCounter(len(words))
	return len(words), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	lines := strings.Count(string(p), "\n")
	*c += LineCounter(lines)
	return lines, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	var name = "Dolly"
	if msg, err := fmt.Fprintf(&c, "hello, %s", name); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
	fmt.Println(c) // "12", = len("hello, Dolly")

	var w WordCounter
	w.Write([]byte("hello"))
	fmt.Println(w) // "1", = len("hello")
	if msg, err := fmt.Fprintf(&w, "hello, %s", name); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
	fmt.Println(w) // "2", = len("hello, Dolly")

}
