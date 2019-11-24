package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webProject/routers"
)

type Config struct {
	Port string
	Debug bool
}

func Run(cfg Config) {
	//初始化gin
	router := gin.Default()
	//注册路由
	routers.RegisterRouters(router)
	//设置监听端口、路由
	srv := &http.Server{
		Addr : ":" + cfg.Port,
		Handler : router,
	}
	//并发监听
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen : %s\n", err)
		}
	}()
	//声明一个信号管道
	quit := make(chan os.Signal)
	//监听收到的信号
	//第一个参数表示接收信号的channel, 第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server")
	//请求超时时间控制
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", err)
	}
}