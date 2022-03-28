package main

import (
	"fmt"
	"runtime"
	"time"
)

type Data struct {
	d [1024 * 100]byte
	o *Data
}

//GC能正确处理 "指针循环引⽤"，但⽆法确定 Finalizer 依赖次序,导致内存泄漏
func finalizertest() {
	var a, b Data
	a.o = &b
	b.o = &a
	runtime.SetFinalizer(&a, func(d *Data) { fmt.Printf("a %p final.\n", d) })
	runtime.SetFinalizer(&b, func(d *Data) { fmt.Printf("b %p final.\n", d) })
}
func main() {
	for {
		finalizertest()
		time.Sleep(time.Millisecond)
	}
}
