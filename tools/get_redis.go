package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	var key, addr string
	flag.StringVar(&addr, "addr", "addrss", "配置文件")
	flag.StringVar(&key, "key", "key", "配置文件")
	var address []string
	flag.Parse()
	address = append(address, addr)
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: address,
		//MaxRedirects: c.config.MaxRetries,
		//ReadOnly:     c.config.ReadOnly,
		//Password:     c.config.Password,
		//MaxRetries:   c.config.MaxRetries,
		//DialTimeout:  c.config.DialTimeout,
		//ReadTimeout:  c.config.ReadTimeout,
		//WriteTimeout: c.config.WriteTimeout,
		//PoolSize:     c.config.PoolSize,
		//MinIdleConns: c.config.MinIdleConns,
		//IdleTimeout:  c.config.IdleTimeout,
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error("redis ping", err)
		os.Exit(1)

	}
	v := rdb.Get(ctx, key)
	fmt.Println("value:", v)
}
