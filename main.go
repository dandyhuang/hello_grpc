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

// Package main implements a stream_server for Greeter service.
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"hello_grpc/dao/mysql/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"hello_grpc/config"
	pb "hello_grpc/service/proto"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest1) (*pb.HelloReply, error) {
	// metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Printf("get metadata error")
	}
	i := 0
	for k, v := range md {

		fmt.Println(i, k, v)
		i = i + 1
	}
	if t, ok := md["key1"]; ok {
		fmt.Printf("timestamp from metadata:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	}
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	proto.JianzhuTool()
	log.Println("=========")
	// dao.GormTool()
	cfg, err := config.Load("./config/hello.yaml")
	if err != nil {
		fmt.Println("cfg err:", err)
		return
	}
	//engine := gin.Default()
	//engine.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//ginAddr := cfg.Server.Addr + ":" + cfg.Server.Port
	//engine.Run(ginAddr)

	lis, err := net.Listen("tcp", ":"+cfg.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listen grpc!")
	// grpc.ServerOption(apply:func() { print("666") })
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
