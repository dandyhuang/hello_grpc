package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu sync.Mutex
	x  int64
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.x++
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := Counter{}
	var state int32
	state = 3
	new := atomic.AddInt32(&state, -1)
	fmt.Println("new:", new)
	var wait sync.WaitGroup
	wait.Add(2)
	for k := 2; k > 0; k-- {
		go func() {
			for i := 2500000; i > 0; i-- {
				c.Inc()
			}
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(c.x)
}
