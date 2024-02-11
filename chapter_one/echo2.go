package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var now time.Time
	for index := range []int{1, 2, 3, 4, 5} {
		now = time.Now()
		s, sep := "", " "
		for _, arg := range os.Args[1:] {
			s += sep + arg
		}
		fmt.Println(index, s, "Time: ", time.Since(now).Seconds()*10000, "Seconds")
		now = time.Now()
		fmt.Println(index, strings.Join(os.Args[1:], " "), "Time: ", time.Since(now).Seconds()*10000, "Seconds")
	}
}
