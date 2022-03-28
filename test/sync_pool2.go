package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"unsafe"
)

func get(a **int) {}
func getPool(pool *sync.Pool) {
	pool.Put(2)
}

func main() {
	count := 0
	pool := sync.Pool{New: func() interface{} {
		count++
		return count
	}}

	//v1 := pool.Get()
	//fmt.Printf("value 1: %v\n", v1)
	pool.Put(10)
	pool.Put(4)
	pool.Put(3)
	// 先取的private的值为10
	fmt.Println("1===:", pool.Get())
	// 取的后进先出，pophead的值，为3
	fmt.Println("2===:", pool.Get())
	// 校验不通的pid值
	go getPool(&pool)

	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	// already full
	pool.Put(11)
	pool.Put(12)
	v2 := pool.Get()
	fmt.Printf("value 2: %v\n", v2)
	// 第一次gc
	runtime.GC()
	v3 := pool.Get()
	fmt.Printf("value 3: %v\n", v3)
	runtime.GC()
	v4 := pool.Get()
	fmt.Printf("value 4: %v\n", v4)
	pool.New = nil
	v5 := pool.Get()
	fmt.Printf("value 5: %v\n", v5)

	f := func(head, tail uint32) uint64 {
		const dequeueBits int = 32
		const mask = 1<<dequeueBits - 1
		return (uint64(head) << dequeueBits) |
			uint64(tail&mask)
	}
	fmt.Println(strconv.FormatInt(int64(f(2, 3)), 2))
	type eface struct {
		typ, val unsafe.Pointer
	}

	vals := make([]eface, 8)
	slot := &vals[0]
	*(*interface{})(unsafe.Pointer(slot)) = 4
	*slot = eface{}
	typ := atomic.LoadPointer(&slot.typ)
	fmt.Println("typeLLLL:", slot)
	if typ == nil {
		// Another goroutine is still cleaning up the tail, so
		// the queue is actually still full.
		fmt.Println("typ is dsdfsss nil")
	}
	*(*interface{})(unsafe.Pointer(slot)) = 5
	fmt.Println(vals)
	value := *(*interface{})(unsafe.Pointer(slot))
	fmt.Println(value)
	*(*interface{})(unsafe.Pointer(slot)) = "string"
	str := *(*interface{})(unsafe.Pointer(slot))
	fmt.Println("str:", str)

	slot.val = nil
	atomic.StorePointer(&slot.typ, nil)
	typ = atomic.LoadPointer(&slot.typ)
	if typ != nil {
		// Another goroutine is still cleaning up the tail, so
		// the queue is actually still full.
		fmt.Println("typ is not nil")
	}

	fmt.Println("runtime.GOMAXPROCS:", runtime.GOMAXPROCS(0), slot)
	s2 := &student{}
	s := student{}
	s.next = new(student)

	s1 := loadPoolChainElt(&s.next)
	s1.sex = 2
	fmt.Println((*unsafe.Pointer)(unsafe.Pointer(s.next)),
		" address:", (*unsafe.Pointer)(unsafe.Pointer(&s.next)), unsafe.Pointer(s1))
	if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&s.next)), unsafe.Pointer(s1), unsafe.Pointer(s2)) {
		fmt.Println("cas ")
	}
}

type student struct {
	pre, next *student
	sex       int
}

func loadPoolChainElt(pp **student) *student {
	return (*student)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(pp))))
}
