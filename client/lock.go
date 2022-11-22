package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var g_trss []string

func test1(args ...string) { //可以接受任意个string参数
	for _, v := range args {
		fmt.Println(v)
	}
	g_trss = args
}
func main2() {
	var strss = []string{
		"qwr",
		"234",
		"yui",
	}
	test1("13", "23")
	test1(strss...)
	fmt.Println("g_trss:", g_trss)
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
