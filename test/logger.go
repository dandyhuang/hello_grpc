package main

import (
	"container/heap"
	"fmt"
	"time"
)

type ss interface {
	name()
}
type int32 int

type student1 struct {
}
type student1ptr *student1

func (receiver string) name() {

}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h IntHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func main() {
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(h)
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)

		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
