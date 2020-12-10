package main

import (
	"fmt"
	"net/http"
	"github.com/golang/sync"

)

//基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
func main() {
	var errG sync.errGroup
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Signal!\n")
	})

	var srv=http.Server{
		Addr:"localhost:8080",
	}


}
