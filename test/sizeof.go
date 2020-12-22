package main

import (
	"fmt"
	"unsafe"
)

type Man struct {
	Name string
	Age  int
	sex  uint32
}

func main() {

	m := Man{Name: "John", Age: 20}

	fmt.Println("name len:", len(m.Name))
	fmt.Println("man size:", unsafe.Sizeof(m))
	fmt.Println("name size:", unsafe.Sizeof(m.Name))
	fmt.Println("age size:", unsafe.Sizeof(m.Age))
	fmt.Println("sex size:", unsafe.Sizeof(m.sex))
}
