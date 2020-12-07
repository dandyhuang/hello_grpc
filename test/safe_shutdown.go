package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

func serverApp(stop <-chan struct{}, index int) error {
	addr := "127.0.0.1:808" + strconv.Itoa(index)
	s := http.Server{Addr: addr}

	go func() {
		<-stop
		fmt.Printf("safe %d shutdown\n", index)
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serverApp(stop, 1)
	}()
	go func() {
		done <- serverApp(stop, 2)
	}()

	var closed bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error:%+v", err)
		}
		if !closed {
			closed = true
			close(stop)
		}
	}
	signal_stop := make(chan os.Signal)
	signal.Notify(signal_stop, os.Interrupt)
	select {
	case <-signal_stop:
	}
}
