package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Config1 struct {
	sync.Mutex
	Name string
}

type Config2 struct {
	*sync.Mutex
	Name string
}

const (
	mutexLocked1 = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShiftx
	mutexWaiterShift = iota
)

func mutex_spin() bool {
	fmt.Println("func spin")
	return true
}

const rwmutexMaxReaders = 1 << 30

func main() {

	old := false
	if old && mutex_spin() || mutex_spin() {
		fmt.Println("spin")
	}

	fmt.Println(mutexLocked1, "}", mutexWoken, "|", mutexStarving, "|", mutexWaiterShiftx)
	var num int32
	num = 0
	atomiNum := atomic.AddInt32(&num, -rwmutexMaxReaders)
	new_num := atomiNum + rwmutexMaxReaders
	fmt.Println(num, " auto:", atomiNum, " new:", new_num)
	c1 := Config1{Name: "1"}
	cc1 := c1
	fmt.Println(cc1.Name)
	cc1.Lock()
	cc1.Unlock()

	c2 := Config2{
		Mutex: &sync.Mutex{},
		Name:  "2",
	}
	cc2 := c2
	fmt.Println(cc2.Name)
	cc2.Lock()
	cc2.Unlock()
}
