package config

import (
	"fmt"
	"goblog/pkg/config"
)

func init() {
	fmt.Println("加载文件: config/session.go 执行 init() 方法 ====== 设置session配置信息")

	config.Add("session", config.StrMap{

		// 目前只支持 Cookie
		"default": config.Env("SESSION_DRIVER", "cookie"),

		// 会话的 Cookie 名称
		"session_name": config.Env("SESSION_NAME", "goblog-session"),
	})
}
