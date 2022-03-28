package main

import (
	"fmt"
	"github.com/valyala/fastrand"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var kinds []uint8
var bases []int

func init() {
	kinds = []uint8{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122}
}

// 随机字符串
func Krand(size int) []byte {
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = kinds[rand.Intn(62)]
	}
	return result
}

func InitBase(nLen, nSlice int) {
	bases = make([]int, nSlice)
	for i := 0; i < nSlice; i++ {
		bases[i] = (nLen / nSlice) * i
		fmt.Println(bases[i])
	}
}

// 随机字符串
func Krand2(nLen, nSlice int) []byte {
	result := make([]byte, nLen)
	w := sync.WaitGroup{}
	w.Add(nSlice)
	for s := 0; s < nSlice; s++ {
		go func(b int) {
			sliceLen := b + nLen/nSlice
			// fmt.Println("slicelen:", sliceLen)
			for i := b; i < sliceLen; i++ {
				result[i] = kinds[rand.Intn(62)]
			}
			w.Done()
		}(bases[s])
	}
	w.Wait()
	return result
}

// 随机字符串
func Krand3(nLen, nSlice int) []byte {
	result := make([]byte, nLen)
	w := sync.WaitGroup{}
	w.Add(nSlice)
	for s := 0; s < nSlice; s++ {
		go func(b int) {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			sliceLen := b + nLen/nSlice
			for i := b; i < sliceLen; i++ {
				result[i] = kinds[r.Intn(62)]
			}
			w.Done()
		}(bases[s])
	}
	w.Wait()
	return result
}

// 随机字符串
func Krand4(nLen, nSlice int) []byte {
	result := make([]byte, nLen)
	w := sync.WaitGroup{}
	w.Add(nSlice)
	for s := 0; s < nSlice; s++ {
		go func(b int) {

			sliceLen := b + nLen/nSlice
			for i := b; i < sliceLen; i++ {
				result[i] = kinds[fastrand.Uint32n(uint32(62))]
			}
			w.Done()
		}(bases[s])
	}
	w.Wait()
	return result
}

func test(x int) (e error) {
	defer func() {
		if e == nil {
			fmt.Println(x)
		}
	}()
	if x > 0 {
		e = nil
		return
	}
	return fmt.Errorf("123")
}

func main() {
	var a int32
	new_a := atomic.AddInt32(&a, 3)
	fmt.Println("new_a:", new_a)

	atomic.AddInt32(&a, 4)
	atomic.LoadInt32(&a)
	fmt.Println("load:", a)
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second)
		close(ch)
	}()
	fmt.Println("1s close:", <-ch)
	nLen := 1024 * 1024
	start := time.Now()
	for i := 0; i < 100; i++ {
		_ = string(Krand(nLen))
	}
	cost := time.Since(start)
	fmt.Println(cost / 100)

	InitBase(nLen, 4)
	start = time.Now()
	for i := 0; i < 100; i++ {
		_ = string(Krand2(nLen, 4))
	}
	cost = time.Since(start)
	fmt.Println(cost / 100)

	start = time.Now()
	for i := 0; i < 100; i++ {
		_ = string(Krand3(nLen, 4))
		if i < 2 {
			// fmt.Println("Krand3", res)
		}
	}
	cost = time.Since(start)
	fmt.Println("krand3", cost/100)

	start = time.Now()
	for i := 0; i < 100; i++ {
		_ = string(Krand4(nLen, 4))
		if i < 2 {
			// fmt.Println("Krand4", res)
		}
	}
	cost = time.Since(start)
	fmt.Println("krand4", cost/100)
}
