package main

import (
	"sync/atomic"
	"time"
	"unsafe"
)

func fib(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return fib(i - 1)
}

func main() {
	demo1()
}

func demo1() {
	var xDemo1 uint64
	xAddr := uintptr(unsafe.Pointer(&xDemo1))
	println("demo1 before stack copy xDemo1 : ", xDemo1, " xDemo1 pointer: ", &xDemo1, xAddr)

	fib(10000000)

	xPointer := (*uint64)(unsafe.Pointer(xAddr))
	atomic.AddUint64(xPointer, 1)
	println("demo1 after stack copy xDemo1 : ", xDemo1, " xDemo1 pointer:", &xDemo1, xAddr, xPointer)
	time.Sleep(time.Second * 2)
	//pDemo1 := (*uint64)(unsafe.Pointer(&xDemo1))
	//atomic.AddUint64(pDemo1, 1)
	//println("demo1 after stack copy xDemo1 : ", xDemo1, " xDemo1 pointer:", &xDemo1)
}
