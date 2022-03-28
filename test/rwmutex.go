package main

import (
	"fmt"
	"sync"
)

type s struct {
	readerCount int32
}

func main() {
	l := new(sync.RWMutex)
	l.RUnlock()
	fmt.Println("1")
	l.RLock()
}
