package bootstrap

import (
	"fmt"
	"goblog/pkg/route"
	"goblog/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	fmt.Println("加载文件: bootstrap/route.go 执行 SetupRoute() 方法 ====== 路由初始化")

	// 创建路由实例
	router := mux.NewRouter()
	// 注册网页路由规则
	routes.RegisterWebRoutes(router)
	routes.RegisterApiRoutes(router)


	// 设置路由实例
	route.SetRoute(router)
	// 返回路由实例
	return router
}
