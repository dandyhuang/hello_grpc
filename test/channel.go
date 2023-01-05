package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

func AcquireLock(lock_key string, lock_value string, timeout uint32) bool {
	// SET lock_key lock_value NX PX ttl_with_seconds
	// return is_success
	return true
}
func ReleaseLock(lock_key string, lock_value string) bool {
	// 在释放锁的时候加入乐观锁校验, 并通过lua脚本保证原子性
	// return_val = eval (
	//           if redis.call("get",lock_key) == lock_value then
	//                return redis.call("del",lock_key)
	//           else
	//                return 0
	//           end
	// )
	return lock_value != ""
}
func Process(lock_key string, lock_value string, timeout uint32) {
	if AcquireLock(lock_key, lock_value, timeout) {
		// 无论业务逻辑执行是否成功, 一定释放锁
		defer func() {
			release_success := ReleaseLock(lock_key, lock_value)
			if !release_success {
				// 如果锁释放失败, 说明锁超时, 其他人已经获取了锁, 需要根据业务决定是否rollback刚刚的操作
				// maybe rollback??
			}
		}()
		// do something
		// maybe process over lock TTL
	}
}

type task func()

type taskPoolSimple struct {
	work chan task     // task channel
	sem  chan struct{} // gr pool size

	wg sync.WaitGroup

	once sync.Once
	done chan struct{}
}

func (p *taskPoolSimple) AddTask(t task) bool {
	select {
	case <-p.done:
		return false
	default:
	}

	select {
	case <-p.done:
		return false
	case p.work <- t:
		fmt.Println("add work")
	case p.sem <- struct{}{}:
		fmt.Println("add sem")
		p.wg.Add(1)
		go p.worker(t)
	}
	return true
}

func (p *taskPoolSimple) worker(t task) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "%s goroutine panic: %v\n%s\n",
				time.Now(), r, string(debug.Stack()))
		}
		p.wg.Done()
		<-p.sem
	}()
	fmt.Println("work1")
	t()
	fmt.Println("work2")
	for t := range p.work {
		fmt.Println("get work")
		t()
	}
}

func main() {
	pool := &taskPoolSimple{
		work: make(chan task),
		sem:  make(chan struct{}, 1),
		done: make(chan struct{}),
	}
	go pool.AddTask(func() {
		fmt.Println("task1")
		time.Sleep(time.Second * 10)
	})
	go pool.AddTask(func() {
		fmt.Println("task2")
		time.Sleep(time.Second * 20)
	})
	go pool.AddTask(func() {
		fmt.Println("task3")
		time.Sleep(time.Second * 20)
	})
	go pool.AddTask(func() {
		fmt.Println("task4")
		time.Sleep(time.Second * 4)
	})
	time.Sleep(time.Second * 30)
}
