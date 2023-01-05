package main

import (
	"fmt"
)

func main() {
	qq := []interface{}{"DSd", "fds", "666"}
	hh(qq...)
	num := []int{1, 2, 3}
	hh(num)
}

func hh(arg ...any) {
	fmt.Printf("%v,%v\n", arg...)

}
