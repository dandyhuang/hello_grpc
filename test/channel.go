package main

import (
	"fmt"
	"go/types"
	"sort"
	"time"
)

func AcquireLock(lock_key string, lock_value string, timeout uint32) bool {
	// SET lock_key lock_value NX PX ttl_with_seconds
	// return is_success
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
	return return_val != 0
}
func Process(lock_key string, lock_value string, timeout uint32) {
	if AcquireLock(lock_key, lock_value, timeout) {
		// 无论业务逻辑执行是否成功, 一定释放锁
		defer func() {
			release_success := ReleaseLock(lock_key)
			if !release_success {
				// 如果锁释放失败, 说明锁超时, 其他人已经获取了锁, 需要根据业务决定是否rollback刚刚的操作
				// maybe rollback??
			}
		}()
		// do something
		// maybe process over lock TTL
	}
}

func main() {
	ch := make(chan int, 1)
	go func() {
		num := <-ch
		fmt.Println(num)
	}()
	time.Sleep(time.Second * 1)
	ch <- 2
	<-ch
	fmt.Println("sssss")
	time.Sleep(time.Second * 1)
}
