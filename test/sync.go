package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func PK(endpoint, metric string, tags string) string {
	ret := bufferPool.Get().(*bytes.Buffer)
	ret.Reset()
	defer bufferPool.Put(ret)

	ret.WriteString(endpoint)
	ret.WriteString("/")
	ret.WriteString(metric)
	ret.WriteString("/")
	ret.WriteString(tags)
	return ret.String()
}

func main() {
	str := PK("10.62.6.149", "So_sub_dp_mod1_vs2_lr1_2_Count", "project=vre,module=compute,biz_cat=sorting")
	fmt.Println(str)
	endList := make([]string, 0, 3)
	endList = append(endList, "eeeesdafdsa")
	endList = append(endList, "2")
	endList = append(endList, "3")
	fmt.Printf("len=%d cap=%d slice=%v\n", len(endList), cap(endList), endList)
	endList = append(endList, "4")
	fmt.Printf("len=%d cap=%d slice=%v\n", len(endList), cap(endList), endList)
	fmt.Println(endList)
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		res := <-done
		fmt.Println(i, ":", res)
	}
}
