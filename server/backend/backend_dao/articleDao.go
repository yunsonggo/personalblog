package backend_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type BackendArticleDao interface {
	// 获取所有文章
	QueryArticleAll() (articleAll []web_model.ArticleModel,err error)
	// 根据ID 删除文章
	DeleteArticleById(id int64) (err error)
	// 添加文章
	InsertArticle(article web_model.ArticleModel) (err error)
	// 根据ID 获取文章
	QueryArticleById(id int64) (article web_model.ArticleModel,err error)
	// 根据ID 修改文章
	UpdateArticleById(article web_model.ArticleModel) (err error)
	// 根据ID 字段 修改文章图片 视频
	UpdateArticleImageOrVideo(id int64,field,path string) (err error)
}

type backendArticleDao struct {}

func NewBackendArticleDao() BackendArticleDao {
	return &backendArticleDao{}
}

// 获取所有文章
func (bad *backendArticleDao) QueryArticleAll() (articleAll []web_model.ArticleModel,err error) {
	err = dial.DB.Find(&articleAll)
	return
}
// 根据ID 删除文章
func (bad *backendArticleDao) DeleteArticleById(id int64) (err error) {
	article := new(web_model.ArticleModel)
	_,err = dial.DB.Where("id = ?",id).Delete(article)
	return
}
// 添加文章
func (bad *backendArticleDao) InsertArticle(article web_model.ArticleModel) (err error) {
	_,err = dial.DB.InsertOne(&article)
	return
}
// 根据ID 获取文章
func (bad *backendArticleDao) QueryArticleById(id int64) (article web_model.ArticleModel,err error) {
	_,err = dial.DB.Where("id = ?",id).Get(&article)
	return
}
// 根据ID 修改文章
func (bad *backendArticleDao) UpdateArticleById(article web_model.ArticleModel) (err error) {
	_,err = dial.DB.Where("id = ?",article.Id).Cols("state","menu_id","category_id","is_index",
		"is_top","is_bottom","is_original","link_url","author","author_id","title","desc","banner","cover",
		"content","video","keyword").Update(&article)
	return
}
// 根据ID 字段 修改文章图片 视频
func (bad *backendArticleDao) UpdateArticleImageOrVideo(id int64,field,path string) (err error) {
	newArticle := new(web_model.ArticleModel)
	if field == "banner" {
		newArticle.Banner = path
	} else if field == "cover" {
		newArticle.Cover = path
	} else if field == "video" {
		newArticle.Video = path
	}
	_,err = dial.DB.Where("id = ?",id).Update(newArticle)
	return
}