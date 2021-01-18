package controllers

import (
	"fmt"
	"goblog/pkg/view"
	"net/http"
)

// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
//func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
//}

// About 关于我们页面
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("加载文件: app/http/controllers/pages_controller.go 执行 About() 方法 ====== 关于我们")

	view.RenderSimple(w,view.D{},"about.about")
	//fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
	//	"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

// NotFound 404 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("加载文件: app/http/controllers/pages_controller.go 执行 NotFound() 方法 ====== 404页面")
	w.WriteHeader(http.StatusNotFound)
	view.RenderSimple(w,view.D{},"notfound.not_found")

	//fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}
