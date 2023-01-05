package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	for i := 0; i < 100; i++ {
		myrand := random(1000, 9999)
		fmt.Println(myrand)
	}
}
