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
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/metadata"
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
	var addr string
	flag.StringVar(&addr, "addr", "10.193.49.142:19802", "配置文件")
	flag.Parse()
	fmt.Println("addr:", addr)
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
			Imei: "864022038223938",
			ModelName: "hello_world",
		},
		ReqId: "e423sdfs",
	}
	req.Recommend = append(req.Recommend, &rec.RankRecommendInfo{
		SceneId: 836,
	})
	details, err := ptypes.MarshalAny(&req)
	if err != nil {
		panic(err)
	}
	cr:=&rec.CommonRequest{}
	cr.Request = details
	anyName, err := ptypes.AnyMessageName(cr.Request )
	if err != nil {
		log.Println("msg name", err)
	}
	req2:=&rec.RankRecommendRequest{}
	if err = ptypes.UnmarshalAny(cr.Request, req2); err != nil {
		log.Println("test 1 error", err)
	}
	log.Println("anyName:", anyName, "req2", req2)
	r, err := c.Process(ctx, cr)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	foo := &rec.RankRecommendResponse{}
	if err = ptypes.UnmarshalAny(r.Response, foo); err != nil {
		log.Println("error", err)
	}
	log.Printf("Greeting: %v", foo)
	m := jsonpb.Marshaler{}
	result, err := m.MarshalToString(r.Response)
	fmt.Println("res;", result)

	err = proto.Unmarshal(r.Response.Value, foo)
	fmt.Println("value:",err,  foo)
}
