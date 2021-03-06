package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	guard_pb "hello_grpc/proto/guard"
	"log"
)

var data = flag.String("get", "CkoSQDgyOGEzMTk1Y2MyOWIzZWFlY2I2NTE2YWQzNmJlODEzN2QyODFhNTQ1NWUwMTkxMzc5YjkxZDMyZDhiYWIwOTEyBlYxODM2QRIAGhQSBHdpZmkqDDE3Mi4yNy41MS4yOCKXBxIWcXVpY2tfZ2FtZV9kaXN0X21vZHVsZSoFCISzkWQqBQixuZFkKgUIhbORZCoFCP6zkWQqBQiZtpFkKgUIobWRZCoFCK24kWQqBQibsZFkKgUIuc+RZCoFCJaxkWQqBQjxwpFkKgUIlbGRZCoFCIqzkWQqBQjLzJFkKgUIibORZCoFCOLPkWQqBQiOs5FkKgUIhcGRZCoFCIa0kWQqBQilsZFkKgUIlrORZCoFCJu3kWQqBQi5s5FkKgUIo7GRZCoFCJWzkWQqBQiOtJFkKgUIl9aRZCoFCNSxkWQqBQiez5FkKgUIhLGRZCoFCJG0kWQqBQiDsZFkKgUIhbGRZCoFCPCzkWQqBQiJtZFkKgUIvrqRZCoFCICxkWQqBQj5sZFkKgUIgrGRZCoFCIGxkWQqBQifz5FkKgUIz8+RZCoFCIixkWQqBQiHsZFkKgUI+bKRZCoFCOq0kWQqBQjgsZFkKgUIj7GRZCoFCIe3kWQqBQiOsZFkKgUIgLORZCoFCIi3kWQqBQiFt5FkKgUIvc2RZCoFCIqxkWQqBQjisZFkKgUI/LKRZCoFCIS3kWQqBQjzwJFkKgUIkrGRZCoFCIm3kWQqBQiKt5FkKgUIvrGRZCoFCOzNkWQqBQi/sZFkKgUI2LORZCoFCOS2kWQqBQjhspFkKgUIwrGRZCoFCMa1kWQqBQivv5FkKgUIj76RZCoFCMuxkWQqBQjGsZFkKgUI5LORZCoFCKy7kWQqBQjNsZFkKgUI27SRZCoFCJS4kWQqBQj1sZFkKgUIqLGRZCoFCKi2kWQqBQiIwpFkKgUIsbGRZCoFCKyxkWQqBQiusZFkKgUIn7ORZCoFCMazkWQqBQizsZFkKgUIsrGRZCoFCLWxkWQqBQi0sZFkKgUIorORZCoFCLqxkWQqBQi2sZFkKgUIr7KRZCoFCLmxkWQqBQiqs5FkKgUIq76RZGIaCghwb3NpdGlvbhIOcXVpY2tnYW1lLmlkeHhiIQoHYWxnTmFtZRIWcXVpY2tfZ2FtZV9kaXN0X21vZHVsZWIRCgRwYWdlEglxdWlja2dhbWViKwoJcmVxdWVzdElkEh4yMDIwMTExNjE3NDkxNXhrU1M3alczVXpnU1RCVmFiGgoJdGltZXN0YW1wEg0xNjA1NTIwMTU1ODg2igEmGiQKCnRmX3NlcnZpY2USFnF1aWNrX2dhbWVfZGlzdF9tb2R1bGWSAQMzNzk=", "get")

func main() {
	flag.Parse()
	pb, err := base64.StdEncoding.DecodeString(*data)
	if err != nil {
		fmt.Printf("base64 decode failure, error=[%v]\n", err)
	}
	req := &guard_pb.RecomRequest{}
	if err := proto.Unmarshal(pb, req); err != nil {
		log.Fatalln("Failed to parse search request:", err)
	}
	fmt.Printf("req =[%+v] =============== \n", req)

	buf, _ := proto.Marshal(req)
	req1 := &guard_pb.RecomRequest1{}
	if err = proto.Unmarshal(buf, req1); err != nil {
		log.Fatalln("Failed to parse search request1:", err)
	}
	fmt.Printf("req1 = [%+v]\n", req1)
}
