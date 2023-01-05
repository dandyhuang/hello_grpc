package link

import (
	_ "unsafe" // for go:linkname
)

//go:linkname hello hello_grpc/test/utils.helloWorld
func hello() string {
	return "private.hello()"
}
