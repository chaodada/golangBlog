package controllers

import (
	"fmt"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/view"
	"net/http"

	"gorm.io/gorm"
)

// BaseController 基础控制器
type BaseController struct {
}

// ResposeForSQLError 处理 SQL 错误并返回
func (bc BaseController) ResposeForSQLError(w http.ResponseWriter, err error) {
	if err == gorm.ErrRecordNotFound {
		// 3.1 数据未找到
		w.WriteHeader(http.StatusNotFound)
		view.RenderSimple(w,view.D{},"notfound.not_found")
		//fmt.Fprint(w, "404 文章未找到")
	} else {
		// 3.2 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	}
}

// ResposeForUnauthorized 处理未授权的访问
func (bc BaseController) ResposeForUnauthorized(w http.ResponseWriter, r *http.Request) {
	flash.Warning("未授权操作！")
	http.Redirect(w, r, "/", http.StatusFound)
}
