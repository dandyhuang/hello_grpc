package main

import (
	"fmt"
	"hello_grpc/test/utils"
	"sync"
)

func main() {
	fmt.Println(utils.LinkHello())
	sm := sync.Mutex{}
	sm.Lock()
}
