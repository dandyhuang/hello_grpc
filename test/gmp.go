package main

import (
	"log"
	"math"
	"runtime"
	"sort"
	"time"
)

func main() {
	panic("sss")
	defer print("1")
	time.Since()

	runtime.GOMAXPROCS(1) // 限制为单核
	go func() {
		var count = 0
		for {
			if count == math.MaxInt64 {
				count = 0
			}
			count++
			log.Println("count ++")
		}
	}()
	time.Sleep(time.Second)

	go func() {
		for {
			log.Println("call func 2")
		}
	}()
	sort.Slice()
	select {}
}
