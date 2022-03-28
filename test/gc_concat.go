package main

import (
	"fmt"
	"sync"

	//	"github.com/gogf/gf/os/grpool"
	// "github.com/Jeffail/tunny"
	"github.com/panjf2000/ants/v2"
	"os"
	"runtime"
	"runtime/trace"
	"sync/atomic"
	"time"
)

var (
	stop  int32
	count int64
	sum   time.Duration
)

//func concat() {
//	numCPUs := runtime.NumCPU()
//	p := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
//		s := "Go GC"
//		s += " " + "Hello"
//		s += " " + "World"
//		_ = s
//		return s
//	})
//	defer p.Close()
//
//	//defer ants.Release()
//	////wg := sync.WaitGroup{}
//	////p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
//	////	s := "Go GC"
//	////	s += " " + "Hello"
//	////	s += " " + "World"
//	////	_ = s
//	////	wg.Done()
//	////})
//	////defer p.Release()
//	////// Submit tasks one by one.
//	////for i := 0; i < 800; i++ {
//	////	wg.Add(1)
//	////	_ = p.Invoke(int32(i))
//	////}
//	////wg.Wait()
//	//
//	//var wg sync.WaitGroup
//	//syncCalculateSum := func() {
//	//	s := "Go GC"
//	//	s += " " + "Hello"
//	//	s += " " + "World"
//	//	_ = s
//	//	wg.Done()
//	//}
//	//for i := 0; i < 800; i++ {
//	//	wg.Add(1)
//	//	_ = ants.Submit(syncCalculateSum)
//	//}
//	//wg.Wait()
//}

func concatAnts() {
	//defer ants.Release()
	//start := time.Now()
	//var wg sync.WaitGroup
	//syncCalculateSum := func() {
	//	s := "Go GC"
	//	s += " " + "Hello"
	//	s += " " + "World"
	//	_ = s
	//	wg.Done()
	//}
	//for i := 0; i < 800; i++ {
	//	wg.Add(1)
	//	_ = ants.Submit(syncCalculateSum)
	//}
	//wg.Wait()
	//cost := time.Since(start)
	//fmt.Printf("cost ant =[%s]", cost)

	start1 := time.Now()
	count := 0
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		s := "Go GC"
		s += " " + "Hello"
		s += " " + "World"
		_ = s
		count++
		wg.Done()
		// fmt.Println(s)
	}, ants.WithNonblocking(true))
	defer p.Release()
	for i := 0; i < 800; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()
	cost1 := time.Since(start1)
	fmt.Printf("cost1 ant =[%s], count=[%d]\n", cost1, count)
}

//func concat() {
//	gpool := grpool.New(10)
//	//wg := sync.WaitGroup{}
//	for n := 0; n < 800; n++ {
//		//wg.Add(8)
//		//for i := 0; i < 8; i++ {
//		gpool.Add(func() {
//			s := "Go GC"
//			s += " " + "Hello"
//			s += " " + "World"
//			_ = s
//			//wg.Done()
//		})
//		//}
//		//wg.Wait()
//	}
//}

func concat() {
	numCPUs := runtime.NumCPU()
	wg := sync.WaitGroup{}
	for n := 0; n < 100; n++ {
		wg.Add(numCPUs)
		for i := 0; i < numCPUs; i++ {
			go (func() {
				s := "Go GC"
				s += " " + "Hello"
				s += " " + "World"
				_ = s
				wg.Done()
			})()
		}
		wg.Wait()
	}
}

//func concat() {
//	for n := 0; n < 100; n++ {
//		for i := 0; i < 8; i++ {
//			go func() {
//				s := "Go GC"
//				s += " " + "Hello"
//				s += " " + "World"
//				_ = s
//			}()
//		}
//	}
//}contex

type MyError struct {
	Desc string
}

func (myErr MyError) Error() string {
	return myErr.Desc
}

func IsMyError(err error) bool {
	my, ok := err.(MyError)
	return ok && my.Desc == ""
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	go func() {
		var t time.Time
		for atomic.LoadInt32(&stop) == 0 {
			t = time.Now()
			runtime.GC()
			sum += time.Since(t)
			count++
		}
		fmt.Printf("GC spend avg: %v\n", time.Duration(int64(sum)/count))
	}()
	concat()
	atomic.StoreInt32(&stop, 1)
	concatAnts()
}
