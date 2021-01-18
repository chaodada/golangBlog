package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

// RemoveTrailingSlash 除首页以外，移除所有请求路径后面的斜杆
func RemoveTrailingSlash(next http.Handler) http.Handler {

	fmt.Println("加载文件: app/http/middlewares/remove_trailing_slash.go 执行 RemoveTrailingSlash() 方法 ====== 除首页以外，移除所有请求路径后面的斜杠")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// 2. 将请求传递下去
		next.ServeHTTP(w, r)
	})
}
