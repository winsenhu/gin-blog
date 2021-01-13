package main

import (
    "fmt"
    "log"
    "syscall"

    "github.com/fvbock/endless"

    "gin-blog/routers"
    "gin-blog/pkg/setting"
)

// @title gin-blog
// @version 1.0
// @description  Golang api of demo
// @termsOfService http://github.com

// @contact.name API Support
// @contact.url http://www.cnblogs.com
// @contact.email ×××@qq.com

//@host 127.0.0.1:8000
func main() {
    // router := routers.InitRouter()


    // s := &http.Server{
    //     Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
    //     Handler:        router,
    //     ReadTimeout:    setting.ReadTimeout,
    //     WriteTimeout:   setting.WriteTimeout,
    //     MaxHeaderBytes: 1 << 20,
    // }

    // s.ListenAndServe()

    endless.DefaultReadTimeOut = setting.ReadTimeout
    endless.DefaultWriteTimeOut = setting.WriteTimeout
    endless.DefaultMaxHeaderBytes = 1 << 20
    endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

    server := endless.NewServer(endPoint, routers.InitRouter())
    server.BeforeBegin = func(add string) {
        log.Printf("Actual pid is %d", syscall.Getpid())
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Printf("Server err: %v", err)
    }
}