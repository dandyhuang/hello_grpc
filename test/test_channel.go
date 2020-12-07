package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 10)
	ch <- "sdfsd"
	ch <- "ddd"

	go func() {
		for data := range ch {
			fmt.Println("data:", data)
		}
	}()
	time.Sleep(1 * time.Second)
}
