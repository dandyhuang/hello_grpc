package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println("dfdsfd")
	ch:=make(chan int, 1)
	t:=time.Tick(time.Second * 2)


	go func() {
		select {
		case <-ch:
			fmt.Println("3232")
		case <-ctx.Done():
			fmt.Println("werwere")

		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("dfddfdfd")
	ch<-1
	time.Sleep(time.Second * 2)
	/*
	 数据的发送者才能决定这个 channel 什么时候能被关闭。如果有人在读 channel, 有人在写 channel，
	 关闭通道后会有 panic 异常。所以先让 http shutdown, 再调用 tracker 的 shutdown
	*/
}

/*
 引入 channel, 它是一个 worker 的工作模型，只需要启动一两个或者两三个 goroutine，在后台去消费 channel 里面的数据，
 它在背后做上报的工作就好了，就不用大量启动 goroutine, 这样性能会好很多。
*/

func NewTracker() *Tracker {
	return &Tracker{ch: make(chan string, 10)}
}

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		fmt.Println("event========")
		return nil
	case <-ctx.Done():
		return ctx.Err() // Err() 返回一个错误，表示 channel 被关闭的原因。例如是被取消，还是超时。
	}
}

/*
 模拟层上报，一直消费 channel 里面的数据
*/

func (t *Tracker) Run() {
	fmt.Println("444444:  ", t.ch)
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println("输出：", data)

	}
	fmt.Println("dfsdffdsdf")
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop: // 通过 t.stop 知道 Run 函数什么时候退出
		fmt.Println("退出了111")
	case <-ctx.Done():
		fmt.Println("退出了222")
	}
}