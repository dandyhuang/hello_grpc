package main

import "sync"

func main() {
	count := 0
	for {
		x, y, a, b := 0, 0, 0, 0
		count++
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			a = 1
			x = b
			println("thread1 done ", count)
			wg.Done()
		}()
		go func() {
			b = 1
			y = a
			println("thread2 done ", count)
			wg.Done()

		}()
		wg.Wait()
		if x == 0 && y == 0 {
			println("执行次数 ：", count)
			break
		}
	}
}
