package config

import (
	"fmt"
	"goblog/pkg/config"
)

func init() {
	fmt.Println("加载文件: config/pagination.go 执行 init() 方法 ====== 设置pagination配置信息")

	config.Add("pagination", config.StrMap{

		// 默认每页条数
		"perpage": 10,

		// URL 中用以分辨多少页的参数
		"url_query": "p",
	})
}
