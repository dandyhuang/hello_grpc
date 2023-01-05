package main

import "time"

// 但是这个是编译器优化。不是线程可见性
func main() {
	running := true
	go func() {
		println("start thread1")
		count := 1
		for running {
			count++
		}
		println("end thread1: count =", count) // 这句代码永远执行不到为什么？
	}()
	go func() {
		println("start thread2")
		for {
			running = false
		}
	}()
	time.Sleep(time.Hour)
}
