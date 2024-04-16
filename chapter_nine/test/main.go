package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()

	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
	time.Sleep(1 * time.Second)

	var mp map[string]int
	res := mp["a"]
	fmt.Println(res)

	for i := 0; i < 100; i++ {
		go fmt.Print(0)
		fmt.Print(1)
	}

	ch1 := make(chan string)
	ch2 := make(chan string)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	mu := sync.Mutex{}
	i := 0

	go func() {
		ticker := time.Tick(1 * time.Second)
		for {
			select {
			case <-ticker:
				mu.Lock()
				fmt.Println("i:", i)
				mu.Unlock()
			default:
				msg := <-ch1
				ch2 <- msg
				mu.Lock()
				i++
				mu.Unlock()
			}
		}
	}()

	go func() {
		for {
			ch1 <- "hello"
			<-ch2
		}
	}()

	<-sigCh

}
