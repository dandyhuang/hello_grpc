package main

import (
	"errors"
	"fmt"
	"github.com/ossrs/go-oryx-lib/aac"
	perror "github.com/pkg/errors"
	"hello_grpc/test/utils"
	"runtime/debug"
)

func foo() error {
	return errors.New("dfdsf")
}

func test1() error {
	//test2()
	//return errors.Wrapf( "query user %s failed", userId)
	//	err:=fmt.Errorf("operate failed %s", "dsd")
	err := perror.Errorf("requires 7+ but only  bytes")
	//err := perror.New("math: square root of negative number")
	return perror.WithStack(err)

}

func test2() {
	return
	test3()
}

func test3() {
	fmt.Printf("%s", debug.Stack())
	debug.PrintStack()
}

func main() {
	adts, _ := aac.NewADTS()
	if _, _, err := adts.Decode(nil); err != nil {
		fmt.Println(fmt.Sprintf("Decode failed, err is %+v", err))
	}
	fmt.Println(utils.Load("dfsd"))

	fmt.Printf(fmt.Sprintf("%+v", test1()))
}
