package config

import (
	"fmt"
	"goblog/pkg/config"
)

func init() {

	fmt.Println("加载文件: config/upload.go 执行 init() 方法 ====== 设置upload配置信息")

	config.Add("upload", config.StrMap{

		// 文件上传目录
		"path": config.Env("UPLOAD_PATH", "public/uploads/"),



		//// 当前环境，用以区分多环境
		//"env": config.Env("APP_ENV", "production"),
		//
		//// 是否进入调试模式
		//"debug": config.Env("APP_DEBUG", false),
		//
		//// 应用服务端口
		//"port": config.Env("APP_PORT", "3000"),
		//
		//// gorilla/sessions 在 Cookie 中加密数据时使用
		//"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
		//
		//// 用以生成链接
		//"url": config.Env("APP_URL", "http://localhost:3000"),
	})
}
