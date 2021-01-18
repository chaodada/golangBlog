package controllers

import (
	"encoding/json"
	"fmt"
	"goblog/pkg/config"
	"io"
	"os"
	"time"

	"net/http"
)

//type BaseJsonBean struct {
//	Code    int         `json:"code"`
//	Data    interface{} `json:"data,omitempty"`
//	Message string      `json:"message"`
//}

type BaseJsonBean struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error,omitempty"`
	Msg   string `json:"msg,omitempty"`
	//ErrorMessage string  `json:"errorMessage,omitempty"`
	Data interface{} `json:"data,omitempty"`
	//Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

//type ReturnData struct {
//	Code int           `json:"code"`
//	Msg  string        `json:"msg"`
//	Data struct {
//		URL string `json:"url"`
//	} `json:"data,omitempty"`
//	//Data []interface{} `json:"data,omitempty"`
//}

// ApiController 处理静态页面
type ApisController struct {
	BaseController
}

func (ac *ApisController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := NewBaseJsonBean()
	result.Msg = "接口不存在"
	result.Code = 199
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
	return
}

func (ac *ApisController) Api(w http.ResponseWriter, r *http.Request) {
	fmt.Println("加载文件: app/http/controllers/articles_controller.go 执行 Show() 方法 ====== 文章详情")

}

func (ac *ApisController) Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("加载文件: app/http/controllers/articles_controller.go 执行 Show() 方法 ====== 文章详情")
	w.Header().Set("Content-Type", "application/json")
	// 1 计算当前日期
	dirName := time.Now().Format("20060102")
	// 3 拼接上传保存目录
	filePath := config.GetString("upload.path") + dirName + "/"
	// 图片上传完成之后的url 路径
	fileUri := config.GetString("app.url") + "/uploads/" + dirName + "/"
	// 4 判断文件夹是否存在 不存在则创建
	if !Exists(filePath) {
		os.Mkdir(filePath, 0777)
	}
	// 接收图片
	file, handle, err := r.FormFile("image")
	if err != nil {
		result := NewBaseJsonBean()
		result.Error = "接收图片失败"
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
		return
	}
	defer file.Close()
	contentType := handle.Header["Content-Type"][0]
	if contentType == "image/gif" || contentType == "image/jpg" || contentType == "image/png" || contentType == "image/jpeg" {
		f, err := os.OpenFile(filePath+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			result := NewBaseJsonBean()
			result.Error = "服务端保存图片失败"
			bytes, _ := json.Marshal(result)
			fmt.Fprint(w, string(bytes))
			return
		}
		defer f.Close()
		io.Copy(f, file)
		result := NewBaseJsonBean()
		data := make(map[string]string)
		data["filePath"] = fileUri + handle.Filename
		result.Data = data
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
		return
	} else {
		result := NewBaseJsonBean()
		result.Error = "格式不正确"
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
		return
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 获取类型
//func GetFileContentType(file multipart.File) (string, error) {
//	buffer := make([]byte, 512)
//	_, err := file.Read(buffer)
//	if err != nil {
//		return "", err
//	}
//	contentType := http.DetectContentType(buffer)
//	return contentType, nil
//}
