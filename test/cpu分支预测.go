package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"unsafe"
)

const CacheLinePadSize = 64

// CacheLinePad is used to pad structs to avoid false sharing.
type CacheLinePad struct{ _ [CacheLinePadSize]byte }

func main() {
	data := make([]int, 32678)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(256)
	}
	sort.Sort(sort.IntSlice(data)) // Sort和非Sort
	now := time.Now()
	count := 0
	for i := 0; i < len(data); i++ {
		if data[i] > 128 {
			count += data[i]
		}
	}
	end := time.Since(now)
	fmt.Println("time : ", end.Microseconds(), "ms count = ", count)
	fmt.Println(unsafe.Sizeof(CacheLinePad{}))
}
