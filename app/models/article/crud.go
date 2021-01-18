package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.Preload("User").First(&article, id).Preload("Category").Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// GetByUserID 获取全部文章
//func GetByUserID(uid string) ([]Article, error) {
//	var articles []Article
//	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Preload("Category").Find(&articles).Error; err != nil {
//		return articles, err
//	}
//	return articles, nil
//}

// GetByUserID 获取用户全部文章
func GetByUserID(uid string,r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	var articles []Article
	db := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Preload("Category").Find(&articles)
	_pager := pagination.New(r, db, route.Name2URL("users.show","id",uid), perPage)
	// 2. 获取视图数据
	viewData := _pager.Paging()
	_pager.Results(&articles)
	return articles, viewData, nil
}



// GetByCategoryID 获取分类相关的文章
func GetByCategoryID(cid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")
	//fmt.Println(cid)
	//fmt.Println("打印链接" + route.Name2URL("categories.show"))
	//fmt.Println("打印链接" + route.Name2URL("articles.index"))
	//req := map[string]string{
	//	"id":cid,
	//}
	//req :=[]string{"id", "cid"}
	//balance :=[]string{"id",cid}
	_pager := pagination.New(r, db, route.Name2URL("categories.show","id",cid), perPage)


	//_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
	//return nil, viewData, nil

}
