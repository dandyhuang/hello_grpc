package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	// See http://redis.io/topics/cluster-tutorial for instructions
	// how to setup Redis Cluster.
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"product-recall-03-prd4.redis.dba.vivo.lan:12222",
			"product-recall-03-prd3.redis.dba.vivo.lan:12221",
			"product-recall-03-prd2.redis.dba.vivo.lan:12221",
			"product-recall-03-prd1.redis.dba.vivo.lan:12222",
			"product-recall-03-prd0.redis.dba.vivo.lan:12221"},
	})
	stats := client.Ping()
	fmt.Printf("%+v", stats)
	err := client.ReloadState()
	fmt.Println("err:", err)

	pipe := client.Pipeline()
	pipe.HGetAll("dandy1")
	pipe.HGetAll("dandy3")
	pipe.HGetAll("dandy2")
	cmders, err := pipe.Exec()
	if err != nil {
		fmt.Println("err", err)
	}
	for _, cmder := range cmders {
		cmd := cmder.(*redis.StringStringMapCmd)
		strMap, err := cmd.Result()
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("strMap", strMap)
	}

	client.Pipelined(func(pipe redis.Pipeliner) error {
		res := pipe.Get("key")
		fmt.Printf("%+v", res)
		return nil
	})

	//func (c *ClusterClient) Pipelined(fn func(Pipeliner) error) ([]Cmder, error)
	//redisdb.Pipelined()

}
