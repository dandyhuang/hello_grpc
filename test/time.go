package main

import (
	"fmt"
	"time"
)

func main() {
	<-time.Tick(time.Second)
	<-time.After(time.Second)
	<-time.NewTicker(time.Second).C
	<-time.NewTimer(time.Second).C
	time.AfterFunc(time.Second, func() { /*do*/ })

	//创建一个周期性的定时器
	ticker := time.NewTicker(3 * time.Second)
	fmt.Println("当前时间为:", time.Now())

	go func() {
		for {

			//从定时器中获取数据
			t := <-ticker.C
			fmt.Println("当前时间为:", t)

		}
	}()

	for {
		time.Sleep(time.Second * 10)
	}
}
