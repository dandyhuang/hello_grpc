package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	var t int
	var t1 int64
	line := 8
	var s [1024 * 1024][]int64
	// s := make([][]int64, 0)
	for i := 0; i < 1024*1024; i++ {
		for j := 0; j < line; j++ {
			s[i] = append(s[i], int64(0))
		}
	}
	fmt.Printf("%v | %v\n", unsafe.Sizeof(t), unsafe.Sizeof(t1))
	var sum int64
	start := time.Now()
	for i := 0; i < 1024*1024; i++ {
		for j := 0; j < line; j++ {
			sum += s[i][j]
		}
	}
	fmt.Println("costTime1:", time.Since(start))

	start = time.Now()
	for i := 0; i < line; i++ {
		for j := 0; j < 1024*1024; j++ {
			sum += s[j][i]
		}
	}
	fmt.Println("costTime2:", time.Since(start))

	start = time.Now()
	for i := 0; i < 1024*1024; i++ {
		for j := 0; j < line; j++ {
		}
	}
	fmt.Println("costTime3:", time.Since(start))

	start = time.Now()
	for i := 0; i < line; i++ {
		for j := 0; j < 1024*1024; j++ {
		}
	}
	fmt.Println("costTime4:", time.Since(start))

}
