package main

import (
	"fmt"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	c "goblog/pkg/config"
	"net/http"
)

func init() {
	// 初始化配置信息
	//fmt.Println("加载文件: main.go 执行 init() 方法")
	config.Initialize()
}

func main() {

	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	fmt.Println("加载文件: main.go ====== 启动http服务")

	// 启动一个http服务器
	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
