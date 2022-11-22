package main

import (
	"context"
	"google.golang.org/grpc"
	pb "hello_grpc/api/protocol/mixer"
	"hello_grpc/client/consul"
	"log"
	"os"
	"time"
)

const (
	target = "consul://127.0.0.1:8500/helloworld"
)

func main() {
	consul.Init()
	// Set up a connection to the server.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRankServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		r, err := c.Rank(ctx, &pb.RankRequest{ReqId: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.ReqId)
		time.Sleep(time.Second * 2)
	}
}
