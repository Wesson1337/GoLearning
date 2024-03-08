package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) // Количество символов Unicode
	types := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int // Количество длин кодировок UTF-8
	invalid := 0 // Количество некорректных символов UTF-8

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // Возврат руны, байтов, ошибки
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
			case unicode.IsLetter(r):
				types["Letters"]++
			case unicode.IsDigit(r):
				types["Digits"]++
			case unicode.IsSpace(r):
				types["Spaces"]++
			default:
				types["Other"]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntype\tcount\n")
	for i, n := range types {
		fmt.Printf("%v\t%d\n", i, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid)
	}
}