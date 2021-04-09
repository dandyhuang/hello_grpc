package main

import (
	"fmt"
	"sync"
)

type Config1 struct {
	sync.Mutex
	Name string
}

type Config2 struct {
	*sync.Mutex
	Name string
}

func main() {
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
