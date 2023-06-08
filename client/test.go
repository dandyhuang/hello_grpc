// Charcount computes counts of Unicode characters.
package main

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)

type CountIns struct {
	rCount uint32
}

func atomicTest() {
	endPoint := make([]uint32, 0)
	endPoint = append(endPoint, 111)
	endPoint = append(endPoint, 2222)
	endPoint = append(endPoint, 3333)
	ins := CountIns{}
	for {
		i := atomic.LoadUint32(&ins.rCount) % uint32(len(endPoint))
		endpint := endPoint[i]
		atomic.AddUint32(&ins.rCount, 1)
		// ins.rCount++
		fmt.Println(endpint, i, ins.rCount)
		if ins.rCount == 1000 {
			time.Sleep(time.Second * 2)
		}
	}
}

type endPointInfo struct {
}

func cTest(ins *CountIns) {
	endPoint := make([]uint32, 0)
	endPoint = append(endPoint, 111)
	endPoint = append(endPoint, 2222)
	endPoint = append(endPoint, 3333)

	for {
		i := atomic.LoadUint32(&ins.rCount) % uint32(len(endPoint))
		endpint := endPoint[i]
		atomic.AddUint32(&ins.rCount, 1)
		// ins.rCount++
		fmt.Println(endpint, i, ins.rCount)
		if ins.rCount == 1000 {
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	// ins := CountIns{}

	// go cTest(&ins)
	str := "hello,world"
	arr := strings.Split(str, "hello,")
	fmt.Println(arr) // 输出 ["hello" "world"]
	i := 9
	defer func() {
		fmt.Println(i)
	}()
	i = 10
}
