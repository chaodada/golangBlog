package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterApiRoutes(r *mux.Router) {
	fmt.Println("加载文件: routes/api.go 执行 RegisterWebRoutes() 方法 ====== 注册Api相关路由")
	// 404页面


	apis := new(controllers.ApisController)
	//r.NotFoundHandler = http.HandlerFunc(apis.NotFound)

	r.HandleFunc("/api/upload", middlewares.Auth(apis.Upload)).Methods("POST").Name("apis.upload")
	r.HandleFunc("/api/api", apis.Api).Methods("GET").Name("apis.api")

}