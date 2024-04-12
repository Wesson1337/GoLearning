package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func main() {
	var slc []string
	slc = nil

	println(len(slc))
	println(cap(slc))
	fmt.Printf("%v\n", slc)
	test := append(slc, "nil")
	for i := range 5 {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("%s\n", test)

	ctx := context.Background()

	doneCh := make(chan struct{})
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		defer func() { doneCh <- struct{}{} }()
		testGoro(ctx)
	}()

	go cancelAfter(cancel, 5*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("SIGTERM")
	case <-doneCh:
		fmt.Println("done")
		return
	}

}

func testGoro(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context is cancelled, return")
			return
		default:
			fmt.Println("test")
		}

	}
}

func cancelAfter(cancelFunc context.CancelFunc, d time.Duration) {
	<-time.After(d)
	cancelFunc()
	fmt.Println("cancel")
}
