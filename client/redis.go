package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"time"
)

func main() {
	num:=50
	addr:="12"
	flag.StringVar(&addr, "addr", "addrss", "配置文件")
	var addrss []string
	flag.Parse()
	addrss = append(addrss, addr)
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrss,
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
	ctx:=context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error("redis ping", err)
		os.Exit(1)
	}
	start:=time.Now().UnixNano()
	pipeline := rdb.Pipeline()
	result := make([]*redis.StringCmd, 0)
	for i:=0; i< num; i++ {
	 	pipeline.Set(ctx,"dandy"+string(i), i, 100*time.Second)
	}
	_, _ = pipeline.Exec(ctx)
	result = result[:0]
	set:=time.Now().UnixNano()
	for i:=0; i< num; i++ {
		result = append(result, pipeline.Get(ctx,"dandy"+string(i)))
	}
	_, _ = pipeline.Exec(ctx)
	get:=time.Now().UnixNano()
	fmt.Println("set cost:", set -start, get-set)
	fmt.Print(result)

	result = result[:0]
	errstart:=time.Now().UnixNano()
	g, ctx := errgroup.WithContext(context.Background())
	for i:=0; i <  num ; i++{
		g.Go(func() error {
			result = append(result, rdb.Get(ctx, "dandy"+string(i)))
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return
	}
	errend:=time.Now().UnixNano()
	fmt.Println("errcostl:", errend- errstart)
	fmt.Print(result)
}