package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 	http.HandleFunc("/", handle)

	// 	http.ListenAndServe(":8080", nil)
	// }

	// func handle(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := work(ctx, cancel)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("context cancelled")
			return
		case i := <-ch:
			fmt.Println(i)
		}
	}
}

func work(ctx context.Context, cancel context.CancelFunc) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second * 3)
		}
		cancel()
		return
	}()
	return ch
}
