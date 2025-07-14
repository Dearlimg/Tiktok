package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tiktok001/router/routers"
	"tiktok001/setting"
)

func main() {
	//先不使用viper 完了加
	setting.Init()
	r := routers.NewRouter()
	server := &http.Server{
		Addr:           "0.0.0.0:8081",
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("server run success , run on 0.0.0.0:8081")
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		fmt.Println("err:", err)
	case <-quit:
		ctx, cancel := context.WithTimeout(context.Background(), 1)
		defer cancel() //延迟取消上下文
		//上下文超时时间内优雅关机（将未处理完的请求处理完再关闭服务），超过超时时间时退出
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println("server shutdown:", err)
		}
	}

}
