package main

import (
	"container/list"
	"fmt"
)

type LFUCache struct {
	minFreq, capacity int
	keyTable          map[int]*list.List
	freqTable         map[int]*list.List
}

type LinkNode struct {
	key, val, freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{capacity: capacity,
		minFreq:   0,
		keyTable:  make(map[int]*list.List),
		freqTable: make(map[int]*list.List),
	}
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	val := -1
	if l, ok := this.keyTable[key]; !ok {
		return -1
	} else {
		node := l.Front().Value.(LinkNode)
		val = node.val
		if f, fok := this.freqTable[node.freq]; fok {
			f.Remove(l.Front())
			if f.Len() == 0 {
				delete(this.freqTable, node.freq)
				if this.minFreq == node.freq {
					this.minFreq += 1
				}
			}
		}

		// 插入到 freq + 1 中
		if v, ok := this.freqTable[node.freq+1]; ok {
			v.PushFront(LinkNode{key: node.key,
				val:  node.val,
				freq: node.freq + 1})
		} else {
			freqList := list.New()
			freqList.PushFront(LinkNode{key: key, val: node.val, freq: node.freq + 1})
			this.freqTable[node.freq+1] = freqList
		}

		l := list.New()
		l.PushFront(LinkNode{key: key, val: node.val, freq: node.freq + 1})
		this.keyTable[key] = l
	}
	fmt.Println("get value", val)
	return val
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if v, ok := this.keyTable[key]; !ok {
		// 缓存已满，需要进行删除操作
		if len(this.keyTable) == this.capacity {
			fmt.Println("keytable size:", len(this.keyTable), this.minFreq)
			l := this.freqTable[this.minFreq].Back()
			delete(this.keyTable, l.Value.(LinkNode).key)
			this.freqTable[this.minFreq].Remove(this.freqTable[this.minFreq].Back())
			if this.freqTable[this.minFreq].Len() == 0 {
				delete(this.freqTable, this.minFreq)
			}
		}
		// 第一次插入
		l := list.New()
		l.PushFront(LinkNode{key: key, val: value, freq: 1})
		this.freqTable[1] = l
		this.keyTable[key] = l
		this.minFreq = 1
	} else {
		node := v.Front().Value.(LinkNode)
		fmt.Println("node:", node)
		// 删除该次数下的map
		this.freqTable[node.freq].Remove(this.keyTable[key].Front())
		if f, fok := this.freqTable[node.freq]; fok {
			fmt.Println("len:", f.Len(), " freq:", this.minFreq)
			if f.Len() == 0 {
				delete(this.freqTable, node.freq)
				if this.minFreq == node.freq {
					this.minFreq += 1
				}
			}
		}
		// 插入到 freq + 1 中
		if v, ok := this.freqTable[node.freq+1]; ok {
			v.PushFront(LinkNode{key: key,
				val:  value,
				freq: node.freq + 1})
		} else {
			freqList := list.New()
			freqList.PushFront(LinkNode{key: key, val: value, freq: node.freq + 1})
			this.freqTable[node.freq+1] = freqList
		}
		fmt.Println("size:", this.freqTable[node.freq+1].Len())
		// this.freqTable[node.freq+1].PushFront(LinkNode{key: key, val: value, freq: node.freq + 1})
		n := list.New()
		n.PushFront(LinkNode{key: key, val: value, freq: node.freq + 1})
		this.keyTable[key] = n
	}
}
func main() {
	sl := make([]int, 0, 5)
	sl = append(sl, 2)
	s := sl[len(sl):]
	fmt.Printf("%p, %p, %v\n", sl, s, s)

	lfu := Constructor(2)
	lfu.Put(2, 1)
	lfu.Put(1, 1)
	lfu.Get(1)
	lfu.Put(1, 4)
	lfu.Get(2)
	lfu.Put(4, 4)

}
