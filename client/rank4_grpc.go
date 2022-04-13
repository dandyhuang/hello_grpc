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
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/metadata"
	"hello_grpc/proto/rec"
	"log"
	"time"
)

const (
	addr     = "localhost:50050"
)

func main() {
	//conn, err := grpc.Dial("", grpc.WithInsecure(),
	//	grpc.WithBalancer(
	//		grpc.RoundRobin(
	//			grpclb.NewConsulResolver("127.0.0.1:8500", "grpc.health.v1.add"))))

	// Set up a connection to the stream_server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rec.NewCommonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// metadata 使用
	md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	req:=rec.RankRecommendRequest{
		Device:&rec.DeviceInfo{
			Imei: "sdfdsf",
			ModelName: "hello_world",
		},
		ReqId: "e423sdfs",
	}
	req.Recommend = append(req.Recommend, &rec.RankRecommendInfo{
		SceneId: 1001,
	})
	details, err := ptypes.MarshalAny(&req)
	if err != nil {
		panic(err)
	}
	cr:=&rec.CommonRequest{}
	cr.Request = details

	r, err := c.Process(ctx, cr)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r)
}
