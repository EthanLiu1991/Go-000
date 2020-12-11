package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//基于 errgroup 实现一个 http server 的启动和关闭，
//以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
func main() {
	var srv = http.Server{
		Addr: "localhost:8080"}

	errG, ctx := errgroup.WithContext(context.Background())

	errG.Go(func() error {
		fmt.Println("http server start")
		return srv.ListenAndServe()
	})

	errG.Go(func() error {
		<-ctx.Done()
		fmt.Println("http server done")
		srv.Shutdown(context.Background())
		return ctx.Err()
	})

	errG.Go(func() error {
		signals := []os.Signal{os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
		quit := make(chan os.Signal, len(signals))
		signal.Notify(quit, signals...)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("signal done")
				return ctx.Err()
			case <-quit:
				return errors.New("quit signal")
			}
		}
	})

	fmt.Println(errG.Wait())
}
