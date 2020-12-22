package main

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/mennanov/fieldmask-utils"
	pb "hello_world/service/proto"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	var request pb.UpdateUserRequest
	userDst := &pb.User{Sex: 23} // a struct to copy to
	fmt.Println("err")

	mask, err := fieldmask_utils.MaskFromPaths(request.FieldMask.Paths, generator.CamelCase)
	fmt.Println("err", err)
	// handle err...
	fieldmask_utils.StructToStruct(mask, request.User, userDst)

}
