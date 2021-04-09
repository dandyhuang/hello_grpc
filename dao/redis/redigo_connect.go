package main

import (
	"fmt"
	"time"

	red "github.com/gomodule/redigo/redis"
)

var (
	redisPool   *red.Pool
	redisServer = "127.0.0.1:6379"
)

func TestPing() {
	conn := redisPool.Get()
	res, _ := conn.Do("ping")
	fmt.Println(res.(string))
	conn.Close()
}

func newRedisPool(addr string) *red.Pool {
	return &red.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: 60 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (red.Conn, error) {
			return red.Dial("tcp", addr)
		},
	}
}

func TestPoolMaxIdle() {

	go func() {
		fmt.Println("go---1---")
		conn := redisPool.Get()
		fmt.Println("go---1---: Get Conn")
		defer conn.Close()

		res, err := conn.Do("ping")
		if err != nil {
			fmt.Println("go---1---error:", err)
			return
		}
		fmt.Println("go---1---:", res.(string))

		time.Sleep(3 * time.Second)
	}()

	go func() {
		fmt.Println("go---2---")
		conn := redisPool.Get()
		fmt.Println("go---2---: Get Conn")
		defer conn.Close()

		res, err := conn.Do("ping")
		if err != nil {
			fmt.Println("go---2---error:", err)
			return
		}
		fmt.Println("go---2---:", res.(string))

		res, err = conn.Do("get", "hello")
		if err != nil {
			fmt.Println("go---2---error:", err)
			return
		}
		str, _ := red.String(res, err)
		fmt.Println("go---get hello---:", str)
		time.Sleep(3 * time.Second)
	}()

	go func() {
		fmt.Println("go---3---")
		conn := redisPool.Get()
		fmt.Println("go---3---: Get Conn")
		defer conn.Close()

		time.Sleep(5 * time.Second)
		res, err := conn.Do("ping")
		if err != nil {
			fmt.Println("go---3---error:", err)
			return
		}
		fmt.Println("go---3---:", res.(string))

		time.Sleep(3 * time.Second)
	}()
}

func main() {
	redisPool = newRedisPool(redisServer)
	fmt.Println("-------------PoolStats------------")
	fmt.Println("Active/Idle:", redisPool.ActiveCount(), ":", redisPool.IdleCount())

	fmt.Println("-------------TestPing------------")
	TestPing()

	fmt.Println("-------------PoolStats------------")
	fmt.Println("Active/Idle:", redisPool.ActiveCount(), ":", redisPool.IdleCount())

	time.Sleep(2 * time.Second)

	fmt.Println("-------------PoolStats------------")
	fmt.Println("Active/Idle:", redisPool.ActiveCount(), ":", redisPool.IdleCount())

	fmt.Println("-------------TestPoolMaxIdle------------")
	TestPoolMaxIdle()

	time.Sleep(1 * time.Second)
	fmt.Println("-------------PoolStats------------")
	fmt.Println("Active/Idle:", redisPool.ActiveCount(), ":", redisPool.IdleCount())

	time.Sleep(10 * time.Second)

	fmt.Println("-------------PoolStats------------")
	fmt.Println("Active/Idle:", redisPool.ActiveCount(), ":", redisPool.IdleCount())
	select {}
}
