package main

import (
	"flag"
	"fmt"
	"hello_grpc/codec"
	"net"
	"os"
)

var host = flag.String("host", "10.192.10.35", "host")
var port = flag.String("port", "3112", "port")

func main() {


	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	done := make(chan string)
	go handleWrite(conn, done)
	go handleRead(conn, done)
	fmt.Println(<-done)
	fmt.Println(<-done)
}
func handleWrite(conn net.Conn, done chan string) {
	c := codec.ClientCodec{}
	buf, err := c.GetReqbuf()
	if err != nil {
		fmt.Println("getreqbuf err:", err)
	}
	for i := 10; i > 0; i-- {
		_, e := conn.Write(buf)
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
	}
	done <- "Sent"
}
func handleRead(conn net.Conn, done chan string) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	fmt.Println(string(buf[:reqLen-1]))
	done <- "Read"
}
