package main

import (
	"context"
	"fmt"
	"time"
)

var ch = make(chan int)

func getInject() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	case <-ch:
		fmt.Println("ch out")
	}
}

func main() {

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("3333")
		ch <- 1
		fmt.Println("66666666")
	}()
	getInject()
	time.Sleep(time.Second * 5)
	fmt.Println("endllll")
}
