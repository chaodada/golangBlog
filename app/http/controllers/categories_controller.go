package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// CategoriesController 文章分类控制器
type CategoriesController struct {
	BaseController
}

// Create 文章分类创建页面
func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

// Store 保存文章分类
func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {

	// 1. 初始化数据
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	// 2. 表单验证
	errors := requests.ValidateCategoryForm(_category)

	// 3. 检测错误
	if len(errors) == 0 {
		// 创建文章分类
		_category.Create()
		if _category.ID > 0 {
			flash.Success("创建分类成功！")
			//fmt.Fprint(w, "创建成功！")
			//indexURL := route.Name2URL("categories.show", "id", _category.GetStringID())
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")
	}
}

// Show 显示分类下的文章列表
func (cc *CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("加载文件: app/http/controllers/pages_controller.go 执行 Show() 方法 ====== 分类文章列表页")
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)
	// 2. 读取对应的数据
	_category, err := category.Get(id)
	if err != nil {
		// 栏目不存在
		cc.ResposeForSQLError(w, err)
	} else {

		// 3. 获取结果集
		articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 3)
		if err != nil {
			cc.ResposeForSQLError(w, err)

		} else {

			for key, article := range articles {
				articles[key].Body = view.MarkdownToHtml(article.Body)
			}

			// ---  2. 加载模板 ---
			view.Render(w, view.D{
				"Articles":  articles,
				"PagerData": pagerData,
			}, "articles.index", "articles._article_meta")
		}

	}
	//fmt.Println(_category)

}
