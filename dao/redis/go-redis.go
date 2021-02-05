package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	// See http://redis.io/topics/cluster-tutorial for instructions
	// how to setup Redis Cluster.
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"product-recall-03-prd4.redis.dba.vivo.lan:12222",
			"product-recall-03-prd3.redis.dba.vivo.lan:12221",
			"product-recall-03-prd2.redis.dba.vivo.lan:12221",
			"product-recall-03-prd1.redis.dba.vivo.lan:12222",
			"product-recall-03-prd0.redis.dba.vivo.lan:12221"},
	})
	stats := redisdb.Ping()
	fmt.Println(stats)
	err := redisdb.ReloadState()
	fmt.Println("err:", err)

}
