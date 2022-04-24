package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	num:=500
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
	start:=time.Now().UnixNano()/1e6
	pipeline := rdb.Pipeline()
	result := make([]*redis.StringCmd, 0)
	for i:=0; i< num; i++ {
	 	pipeline.Set(ctx, "dandy"+ strconv.Itoa(i), i, 100*time.Second)
	}
	_, _ = pipeline.Exec(ctx)
	result = result[:0]
	pipeline.Discard()
	set:=time.Now().UnixNano()/1e6
	for i:=0; i< num; i++ {
		result = append(result, pipeline.Get(ctx,"dandy"+ strconv.Itoa(i)))
	}
	cmds, err := pipeline.Exec(ctx)
	get:=time.Now().UnixNano()/1e6
	fmt.Println("get cost:", set -start, get-set)
	fmt.Print(result, "cmds:", cmds)

	for _, v:= range cmds {
		value, err :=v.(*redis.StringCmd).Result()

		fmt.Println("value: ", v.(*redis.StringCmd).String(), " | ",  v.String(), value, err )
	}

	pipeline.Discard()
	result = result[:0]
	errstart:=time.Now().UnixNano()/1e6
	g, ctx := errgroup.WithContext(context.Background())
	var res atomic.Value
	res.Store(make([]*redis.StringCmd, 0))
	for i:=0; i <  num ; i++{
		g.Go(func() error {
			result = append(result,rdb.Get(ctx, "dandy"+ strconv.Itoa(i)))
			// res.Store(rdb.Get(ctx, "dandy"+ strconv.Itoa(i)))
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return
	}
	errend:=time.Now().UnixNano()/1e6
	fmt.Println("errcostl:", errend- errstart)
	fmt.Print(result)
}