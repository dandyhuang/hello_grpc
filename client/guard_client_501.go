package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	guard_pb "hello_grpc/proto/guard"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	flag.Parse()
	var data = flag.String("get", "CkoSQDgyOGEzMTk1Y2MyOWIzZWFlY2I2NTE2YWQzNmJlODEzN2QyODFhNTQ1NWUwMTkxMzc5YjkxZDMyZDhiYWIwOTEyBlYxODM2QRIAGhQSBHdpZmkqDDE3Mi4yNy41MS4yOCKXBxIWcXVpY2tfZ2FtZV9kaXN0X21vZHVsZSoFCISzkWQqBQixuZFkKgUIhbORZCoFCP6zkWQqBQiZtpFkKgUIobWRZCoFCK24kWQqBQibsZFkKgUIuc+RZCoFCJaxkWQqBQjxwpFkKgUIlbGRZCoFCIqzkWQqBQjLzJFkKgUIibORZCoFCOLPkWQqBQiOs5FkKgUIhcGRZCoFCIa0kWQqBQilsZFkKgUIlrORZCoFCJu3kWQqBQi5s5FkKgUIo7GRZCoFCJWzkWQqBQiOtJFkKgUIl9aRZCoFCNSxkWQqBQiez5FkKgUIhLGRZCoFCJG0kWQqBQiDsZFkKgUIhbGRZCoFCPCzkWQqBQiJtZFkKgUIvrqRZCoFCICxkWQqBQj5sZFkKgUIgrGRZCoFCIGxkWQqBQifz5FkKgUIz8+RZCoFCIixkWQqBQiHsZFkKgUI+bKRZCoFCOq0kWQqBQjgsZFkKgUIj7GRZCoFCIe3kWQqBQiOsZFkKgUIgLORZCoFCIi3kWQqBQiFt5FkKgUIvc2RZCoFCIqxkWQqBQjisZFkKgUI/LKRZCoFCIS3kWQqBQjzwJFkKgUIkrGRZCoFCIm3kWQqBQiKt5FkKgUIvrGRZCoFCOzNkWQqBQi/sZFkKgUI2LORZCoFCOS2kWQqBQjhspFkKgUIwrGRZCoFCMa1kWQqBQivv5FkKgUIj76RZCoFCMuxkWQqBQjGsZFkKgUI5LORZCoFCKy7kWQqBQjNsZFkKgUI27SRZCoFCJS4kWQqBQj1sZFkKgUIqLGRZCoFCKi2kWQqBQiIwpFkKgUIsbGRZCoFCKyxkWQqBQiusZFkKgUIn7ORZCoFCMazkWQqBQizsZFkKgUIsrGRZCoFCLWxkWQqBQi0sZFkKgUIorORZCoFCLqxkWQqBQi2sZFkKgUIr7KRZCoFCLmxkWQqBQiqs5FkKgUIq76RZGIaCghwb3NpdGlvbhIOcXVpY2tnYW1lLmlkeHhiIQoHYWxnTmFtZRIWcXVpY2tfZ2FtZV9kaXN0X21vZHVsZWIRCgRwYWdlEglxdWlja2dhbWViKwoJcmVxdWVzdElkEh4yMDIwMTExNjE3NDkxNXhrU1M3alczVXpnU1RCVmFiGgoJdGltZXN0YW1wEg0xNjA1NTIwMTU1ODg2igEmGiQKCnRmX3NlcnZpY2USFnF1aWNrX2dhbWVfZGlzdF9tb2R1bGWSAQMzNzk=", "get")
	pb, err := base64.StdEncoding.DecodeString(*data)
	if err != nil {
		fmt.Printf("base64 decode failure, error=[%v]\n", err)
	}
	//请求客户端
	client := &http.Client{}

	req, e := http.NewRequest("POST", "http://10.193.49.142:8889/recommend/predict?sceneid=813",
		strings.NewReader(*data))
	if e != nil {
		log.Fatalln(e.Error())
	}

	req.Header.Set("Content-Type", "application/x-protobuf")

	response, er := client.Do(req)
	if er != nil {
		log.Fatalln(er.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	gd_req := &guard_pb.RecomRequest{}
	if err := proto.Unmarshal(pb, gd_req); err != nil {
		log.Fatalln("Failed to parse search request:", err)
	}
	fmt.Printf("req =[%+v] =============== \n", gd_req)

	buf, _ := proto.Marshal(gd_req)
	req1 := &guard_pb.RecomRequest1{}
	if err = proto.Unmarshal(buf, req1); err != nil {
		log.Fatalln("Failed to parse search request1:", err)
	}
	fmt.Printf("req1 = [%+v]\n", req1)
}
