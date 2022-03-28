package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"hello_grpc/test/proto/gen"
	"reflect"
)

type Person interface {
	GetAge() int
	SetAge(int)
}

type Man struct {
	Name string
	Age  int
}

func (s Man) GetAge() int {
	return s.Age
}

func (s *Man) SetAge(age int) {
	s.Age = age
}

func f(p Person) {
	p.SetAge(10)
	fmt.Println(p.GetAge())
}

func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

type MyInterface interface {
	Print()
}

// go tool compile -S -N -l interface.go
func main() {
	p := Man{}
	// 不想允许&p
	f(&p)
	var x interface{} = nil

	var y *int = nil
	interfaceIsNil(x)
	interfaceIsNil(y)
	getType(y)
	x = 5
	getInterfaceReflectType(x)
	getInterPBEnum()
}

func getInterPBEnum() {
	jsonStatus := `{
	"userStatus":"ONLINE"
}
`
	req := &gen.HelloRequest{}
	json.Unmarshal([]byte(jsonStatus), req)
	fmt.Println("1111", req)

	protojson.Unmarshal([]byte(jsonStatus), req)
	fmt.Println("ssss", req)
}

func getInterfaceReflectType(a interface{}) {
	typename := reflect.TypeOf(a)
	fmt.Println(typename)
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case *int:
		fmt.Println("the type of a is *int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}
