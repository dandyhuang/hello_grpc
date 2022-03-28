package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					adone = true
					fmt.Println("adone")
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					bdone = true
					fmt.Println("bdone")
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func asChan(vs ...int) <-chan int {
	c := make(chan int, 1)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
func deadlockRange() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	// block c
	for v := range c {
		fmt.Println(v)
	}
}
func sendBlockWhenNil() {
	chUnCache := make(chan float64, 100)
	go func() {
		for i := 0; i < 100; i++ {
			chUnCache <- 1
			fmt.Println("1")
			time.Sleep(1 * time.Second)
		}

	}()
	// 发送4次之后，阻塞发送进程
	time.Sleep(4 * time.Second)
	chUnCache = nil
	go func() {
		select {
		case <-chUnCache:
			fmt.Println("receive")
		}
	}()
	// panic: close of nil channel
	close(chUnCache)
	time.Sleep(3 * time.Second)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
		}
	}
}

func timeOut() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//quit <- 0
	}()
	fibonacci(c, quit)
}
func consumeByRange() {
	// 增加一个调度，不然gmp中没有调度，就会报错
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		// 只有close关掉后， range才会结束， 否则会一直阻塞
		// close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}
func selectUnBlock() {
	c := make(chan int, 100)

	go func() {
		time.Sleep(5 * time.Second)

		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		println("start----for")
		select {
		case i := <-c:
			fmt.Println(i)
		default:
			fmt.Println("default")
		}
		println("start----end")
	}

}

func sendNoBlockWhenNil() {
	chUnCache := make(chan float64, 100)
	go func() {
		for i := 0; i < 100; i++ {
			chUnCache <- 1
			fmt.Println("1")
			time.Sleep(1 * time.Second)
		}

	}()
	// 发送4次之后，阻塞发送进程
	time.Sleep(4 * time.Second)
	// chUnCache = nil
	go func() {
		for {
			select {
			case <-chUnCache:
				fmt.Println("receive")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	// panic: close of nil channel
	close(chUnCache)
	time.Sleep(3 * time.Second)
}

func main() {
	// deadlockRange()
	// sendNoBlockWhenNil()
	// timeOut()
	// consumeByRange()
	// selectUnBlock()
	var group sync.WaitGroup
	group.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer group.Done()
		}()
		group.Wait()
	}

}
