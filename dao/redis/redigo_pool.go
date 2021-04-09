package main

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

var pool *redigo.Pool

func ExecPool(cmd string, key interface{}, args ...interface{}) (interface{}, error) {

	con := pool.Get()
	// connct
	if err := con.Err(); err != nil {
		return nil, err
	}

	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)

	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
	//return nil, nil
}

func main() {
	var addr = "127.0.0.1:6379"
	var password = ""

	pool = PoolInitRedis(addr, password)
	ExecPool("get", "dandy")
	ExecPool("get", "dandy")
	c1 := pool.Get()
	c1.Close()
	c2 := pool.Get()
	c2.Close()

	ExecPool("get", "dandy")
	ExecPool("get", "dandy")
	c3 := pool.Get()
	c4 := pool.Get()
	c5 := pool.Get()
	fmt.Println(c1, c2, c3, c4, c5)
	time.Sleep(time.Second * 3) //redis一共有多少个连接？？

	c3.Close()
	c4.Close()
	c5.Close()
	time.Sleep(time.Second * 3) //redis一共有多少个连接？？

	//下次是怎么取出来的？？
	b1 := pool.Get()
	b2 := pool.Get()
	b3 := pool.Get()
	fmt.Println(b1, b2, b3)
	time.Sleep(time.Second * 5)
	b1.Close()
	b2.Close()
	b3.Close()

	//redis目前一共有多少个连接？？
	for {
		fmt.Println("主程序运行中....")
		time.Sleep(time.Second * 1)
	}
}

// redis pool
func PoolInitRedis(server string, password string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   0, //最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
