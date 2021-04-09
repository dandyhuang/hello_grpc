package main

import (
	"fmt"
	red "github.com/gomodule/redigo/redis"
	"time"
)

type Redis struct {
	pool *red.Pool
}

func NewRedis(pool *red.Pool) *Redis {
	return &Redis{pool: pool}
}

var redis *Redis

func Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {

	con := redis.pool.Get()
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

func initRedis() {
	redis = new(Redis)
	redis.pool = &red.Pool{
		//Wait:        true,
		MaxIdle:     2, //空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   0, //最大数
		Dial: func() (red.Conn, error) {
			c, err := red.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c red.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	//redis.pool = &red.Pool{
	//	MaxIdle:     2,
	//	IdleTimeout: 240 * time.Second,
	//	MaxActive:   3, //0设置为无限制，但是每次都会连接，因为会先关闭
	//	//IdleTimeout: time.Duration(120),
	//	Dial: func() (red.Conn, error) {
	//		return red.Dial(
	//			"tcp",
	//			"127.0.0.1:6379",
	//			//red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
	//			//red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
	//			//red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
	//			//red.DialDatabase(0),
	//			//red.DialPassword("888"),
	//		)
	//	},
	//}
}

func main() {

	println("dsd")
	for i := 0; i < 2 && i > 100; i++ {
		println("dsd")
	}

	initRedis()
	con := redis.pool.Get()
	con1 := redis.pool.Get()
	con.Close()
	con1.Close()
	fmt.Println("init start")
	Exec("set", "hello", "world")
	Exec("set", "dandy", "world")

	fmt.Println(2)
	result, err := Exec("get", "hello")
	if err != nil {
		fmt.Print(err.Error())
	}
	str, _ := red.String(result, err)
	fmt.Print(str)
	time.Sleep(time.Second * 10)
	redis.pool.Close()
}
