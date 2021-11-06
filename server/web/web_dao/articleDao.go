package web_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type ArticleDao interface {
	// 获取某一大类文章(主要用于第一个大类的文章 最先展示)
	QueryIndexArticle(menuId int64,count,start int) (articles []web_model.ArticleModel,err error)
	// 获取各大类首页展示文章
	QueryMenuIndexArticle(menuId int64,count,start int) (articles []web_model.ArticleModel,err error)
	// 获取指定大类和小类文章
	QueryMenuCategoryArticle(menuId,categoryId int64,count,start int) (articles []web_model.ArticleModel,err error)
	// 获取指定大类无小类文章
	QueryMenuArticle(menuId int64,count,start int) (articles []web_model.ArticleModel,err error)
	// 根据	ID 获取文章
	QueryArticleById(articleId int64) (article web_model.ArticleModel,err error)
	// 根据ID 更新 文章某一字段
	UpdateArticleOnlyValueById(articleId int64,fieldName string,newArticle *web_model.ArticleModel) (err error)
	// 根据关键字 搜索文章
	QueryArticleByKeyword(key,keyword string,count,start int) (articleList []web_model.ArticleModel,err error)
}

type articleDao struct {}

func NewArticleDao() ArticleDao {
	return &articleDao{}
}
// 获取某一大类文章(主要用于第一个大类的文章 最先展示)
func (ad *articleDao) QueryIndexArticle(menuId int64,count,start int) (articles []web_model.ArticleModel,err error) {
	err = dial.DB.Where("state = 0 and menu_id = ?",menuId).Limit(count,start).OrderBy("update_time desc").Find(&articles)
	return
}
// 获取各大类首页展示文章
func (ad *articleDao) QueryMenuIndexArticle(menuId int64,count,start int) (articles []web_model.ArticleModel,err error) {
	err = dial.DB.Where("state = 0 and menu_id = ? and is_index = 1",menuId).Limit(count,start).OrderBy("update_time desc").Find(&articles)
	return
}
// 获取指定大类和小类文章
func (ad *articleDao) QueryMenuCategoryArticle(menuId,categoryId int64,count,start int) (articles []web_model.ArticleModel,err error) {
	err = dial.DB.Where("state = 0 and menu_id = ? and category_id = ?",menuId,categoryId).Limit(count,start).OrderBy("update_time desc").Find(&articles)
	return
}
// 获取指定大类无小类文章
func (ad *articleDao) QueryMenuArticle(menuId int64,count,start int) (articles []web_model.ArticleModel,err error) {
	err = dial.DB.Where("state = 0 and menu_id = ?",menuId).Limit(count,start).OrderBy("update_time desc").Find(&articles)
	return
}
// 根据	ID 获取文章
func (ad *articleDao) QueryArticleById(articleId int64) (article web_model.ArticleModel,err error) {
	_,err = dial.DB.Where("state = 0 and id = ?",articleId).Get(&article)
	return
}
// 根据ID 更新 文章某一字段
func (ad *articleDao) UpdateArticleOnlyValueById(articleId int64,fieldName string,newArticle *web_model.ArticleModel) (err error) {
	_,err = dial.DB.Where("state = 0 and id = ?",articleId).Cols(fieldName).Update(newArticle)
	return
}
// 根据关键字 搜索文章
func (ad *articleDao) QueryArticleByKeyword(key,keyword string,count,start int) (articleList []web_model.ArticleModel,err error) {
	err = dial.DB.Where(key + " like ? and state = 0","%"+keyword+"%").Limit(count,start).OrderBy("update_time desc").Find(&articleList)
	return
}
