package main

import (
	"context"
	"google.golang.org/grpc"
	pb "hello_grpc/api/protocol/mixer"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedRankServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.RankRequest) (*pb.RankResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.RankResponse{Version: in.Version}, nil
}

func main()  {
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRankServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
