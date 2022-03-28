package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，
// 以及 linux signal 信号的注册和处理，
// 要保证能够一个退出，全部注销退出。

func main() {

	runtimeC := make(chan struct{})
	g, ctx := errgroup.WithContext(context.Background())
	fmt.Println("get pid:", os.Getpid())
	// http服务
	g.Go(func() error {
		r := gin.Default()
		r.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success1": true,
			})
			runtimeC <- struct{}{}
		})
		server := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}
		go func() {
			<-ctx.Done()
			fmt.Println("http service closed.")
			server.Shutdown(ctx) //关闭服务
			runtimeC <- struct{}{}
		}()
		server.ListenAndServe()
		return nil
	})

	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		select {
		case <-ctx.Done():
			fmt.Println("signal ctx done")
			signal.Stop(sig)
			return nil
		case <-sig:
			return errors.New("close")
		}
	})

	g.Go(func() error {
		<-runtimeC
		return errors.New("close")
	})

	g.Wait() //等待第一个error的返回
	time.Sleep(time.Second * 3)
}
