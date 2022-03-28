package main

import (
	"fmt"
	"time"
)

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}
func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo1 val = %d\n", x)
	}
}
func main() {
	c := make(chan int, 3)
	//	quit := make(chan int, 3)
	x := 5

	go func() {
		time.Sleep(100000000000)
		c <- 5
	}()

	for {
		select {
		case i := <-c:
			fmt.Println("i:", i)

			fmt.Println("quit")
		}
	}

	// Q1第一组实验
	f1 := foo1(&x)
	f2 := foo2(x)
	f1()
	f2()
	f1()
	f2()
	// Q1第二组
	x = 233
	f1()
	f2()
	f1()
	f2()
	// Q1第三组
	foo1(&x)()
	foo2(x)()
	foo1(&x)()
	foo2(x)()
}
