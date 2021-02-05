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

	// method 2
	client.WrapProcessPipeline(func(old func([]redis.Cmder) error) func([]redis.Cmder) error {
		return func(cmds []redis.Cmder) error {
			fmt.Printf("pipeline starting processing: %v\n", cmds)
			err := old(cmds)
			fmt.Printf("pipeline finished processing: %v\n", cmds)
			return err
		}
	})

	client.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Get("key")
		pipe.Get("key1")
		return nil
	})

	//func (c *ClusterClient) Pipelined(fn func(Pipeliner) error) ([]Cmder, error)
	//redisdb.Pipelined()

}
