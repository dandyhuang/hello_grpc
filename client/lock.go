package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
	_ "net/http/pprof"
)

func main2() {
	total := 0
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("total: ", total)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		fmt.Println("3223")
		mutex.Lock()
		defer mutex.Unlock()
		go func() {
			total += 1
		}()
	}
	fmt.Println("1111")
	mutex.Lock()
	defer mutex.Unlock()
	for j := 0; j < 5; j++ {
		go func() {
			total += 1
		}()
	}

}
func main() {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	go http.ListenAndServe("localhost:6060", nil)
	main2()
}