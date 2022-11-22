package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func(ch chan int) {
		fmt.Println("print1111")
		time.Sleep(time.Second * 2)
		fmt.Println("print1111")
		close(ch)
	}(ch)
	ctx := context.Background()
	select {
	case <-ch:
	case <-ctx.Done():
	}
	fmt.Println("1134")
	var secret string
	fmt.Scanln(&secret)
	if secret == "" {
		fmt.Println("")
		return
	}
	var left, right int
	if len(secret)%2 == 0 {
		left = len(secret) / 2
		right = len(secret) / 2
	} else {
		left = len(secret) / 2
		right = len(secret)/2 + 1
	}

	for {
		if left <= 0 || right >= len(secret) {
			break
		}

		if secret[0:left] == secret[right:] {
			break
		}

		left--
		right++
	}
	fmt.Println(secret[0:left])
}
