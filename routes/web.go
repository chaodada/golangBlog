package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	fmt.Println("加载文件: routes/web.go 执行 RegisterWebRoutes() 方法 ====== 注册网页相关路由")

	// 实例化静态页面
	pc := new(controllers.PagesController)
	// 404页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	// 关于我们页面
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")




	// 实例化文章相关页面
	ac := new(controllers.ArticlesController)
	// 主页
	r.HandleFunc("/", ac.Index).Methods("GET").Name("home")
	// 文章详情页
	r.HandleFunc("/articles/{id:[0-9]+}.html", ac.Show).Methods("GET").Name("articles.show")
	// 文章列表页
	r.HandleFunc("/articles.html", ac.Index).Methods("GET").Name("articles.index")
	// 文章创建页
	r.HandleFunc("/articles/create.html", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")
	// 文章创建
	r.HandleFunc("/articles.html", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	// 文章更新页
	r.HandleFunc("/articles/{id:[0-9]+}/edit.html", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	// 文章更新
	r.HandleFunc("/articles/{id:[0-9]+}.html", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")
	// 文章删除
	r.HandleFunc("/articles/{id:[0-9]+}/delete.html", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")


	// 文章分类
	cc := new(controllers.CategoriesController)
	r.HandleFunc("/categories/create.html", middlewares.Auth(cc.Create)).Methods("GET").Name("categories.create")
	r.HandleFunc("/categories.html", middlewares.Auth(cc.Store)).Methods("POST").Name("categories.store")

	r.HandleFunc("/categories/{id:[0-9]+}.html", cc.Show).Methods("GET").Name("categories.show")


	// 用户认证
	auc := new(controllers.AuthController)
	// 注册页面
	r.HandleFunc("/auth/register.html", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	// 注册
	r.HandleFunc("/auth/do-register.html", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")
	// 登陆页面
	r.HandleFunc("/auth/login.html", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	// 登陆
	r.HandleFunc("/auth/dologin.html", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	// 退出登陆
	r.HandleFunc("/auth/logout.html", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

	// 用户信息
	uc := new(controllers.UserController)
	r.HandleFunc("/users/{id:[0-9]+}.html", uc.Show).Methods("GET").Name("users.show")

	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/uploads/").Handler(http.FileServer(http.Dir("./public")))


	// --- 全局中间件 ---

	// 开始会话
	r.Use(middlewares.StartSession)
}
