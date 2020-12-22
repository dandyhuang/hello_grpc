package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

type Man struct {
	Name  string
	Age   int
	sex   uint32
	types uint8
}

func main() {
	fmt.Println("===========以下通过Write把swift写入Learning缓冲器尾部=========")
	newBytes := []byte("swift")
	//创建一个内容Learning的缓冲器
	buf := bytes.NewBuffer([]byte("Learning"))
	//打印为Learning
	fmt.Println(buf.String(), len(newBytes))
	//将newBytes这个slice写到buf的尾部
	buf.Write(newBytes)
	fmt.Println(buf.String())

	m := Man{Name: "John", Age: 20}
	x := "text"
	xBytes := []byte(x)
	var trace_id []byte
	copy(trace_id, "tracechage")
	data := []byte("hello")
	fmt.Println("trade_id:", trace_id)
	fmt.Println("xBytes：", xBytes)
	fmt.Println("data:", string(data))
	fmt.Println("name len:", len(m.Name))
	fmt.Println("man size:", unsafe.Sizeof(m))
	fmt.Println("name size:", unsafe.Sizeof(m.Name))
	fmt.Println("age size:", unsafe.Sizeof(m.Age))
	fmt.Println("sex size:", unsafe.Sizeof(m.sex))
	fmt.Println("types size:", unsafe.Sizeof(m.types))
}
