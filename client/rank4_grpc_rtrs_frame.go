/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/metadata"
	"hello_grpc/api/protocol/rank2"
	"hello_grpc/proto/rec"
	"log"
	"time"
)

func main() {
	//conn, err := grpc.Dial("", grpc.WithInsecure(),
	//	grpc.WithBalancer(
	//		grpc.RoundRobin(
	//			grpclb.NewConsulResolver("127.0.0.1:8500", "grpc.health.v1.add"))))

	// Set up a connection to the stream_server.
	var addr, imei, flagconf string
	flag.StringVar(&addr, "addr", "10.193.34.19:8880", "配置文件")
	flag.StringVar(&imei, "imei", "864022038223938", "配置文件")
	flag.StringVar(&flagconf, "conf", "user1.json", "配置文件")

	flag.Parse()
	fmt.Println("addr:", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), // grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rec.NewCommonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// metadata 使用
	md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	c1 := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c1.Close()

	if err := c1.Load(); err != nil {
		panic(err)
	}
	req := rank2.RecomRequest{}
	if err := c1.Scan(&req); err != nil {
		panic(err)
	}

	log.Println("req:", req)

	details, err := ptypes.MarshalAny(&req)
	if err != nil {
		panic(err)
	}
	cr := &rec.CommonRequest{}
	cr.Request = details
	anyName, err := ptypes.AnyMessageName(cr.Request)
	if err != nil {
		log.Println("msg name", err)
	}
	cr.Reserved2 = uint64(req.Recommend.SceneId)

	log.Println("anyName:", anyName)
	r, err := c.Process(ctx, cr)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	foo := &rank2.RecomResponse{}

	err = proto.Unmarshal(r.Response.Value, foo)
	fmt.Println("value:", err, foo)
}
