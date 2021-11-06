package backend_service

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type BackendArticleServer interface {
	// 获取所有文章
	FindArticleAll() (articleAll []web_model.ArticleModel,err error)
	// 根据ID删除文章
	RemoveArticleById(id int64) (err error)
	// 批量删除文章
	RemoveArticleArrById(idArr []int64) (err error)
	// 添加文章
	AddArticle(article web_model.ArticleModel) (err error)
	// 根据ID 获取文章
	GetArticleById(id int64) (article web_model.ArticleModel,err error)
	// 根据ID 修改文章
	EditArticleById(article web_model.ArticleModel) (err error)
	// 根据ID 字段 修改文章图片 视频
	EditArticleImageOrVideo(id int64,field,path string) (err error)
}

type backendArticleServer struct {}

func NewBackendArticleServer() BackendArticleServer {
	return &backendArticleServer{}
}

var bad = backend_dao.NewBackendArticleDao()

// 获取所有文章
func (bas *backendArticleServer) FindArticleAll() (articleAll []web_model.ArticleModel,err error) {
	return bad.QueryArticleAll()
}
// 根据ID删除文章
func (bas *backendArticleServer) RemoveArticleById(id int64) (err error) {
	return bad.DeleteArticleById(id)
}
// 批量删除文章
func (bas *backendArticleServer) RemoveArticleArrById(idArr []int64) (err error) {
	for _,id := range idArr {
		err = bad.DeleteArticleById(id)
		if err != nil {
			return
		}
	}
	return
}
// 添加文章
func (bas *backendArticleServer) AddArticle(article web_model.ArticleModel) (err error) {
	return bad.InsertArticle(article)
}
// 根据ID 获取文章
func (bas *backendArticleServer) GetArticleById(id int64) (article web_model.ArticleModel,err error) {
	return bad.QueryArticleById(id)
}
// 根据ID 修改文章
func (bas *backendArticleServer) EditArticleById(article web_model.ArticleModel) (err error) {
	return bad.UpdateArticleById(article)
}
// 根据ID 字段 修改文章图片 视频
func (bas *backendArticleServer) EditArticleImageOrVideo(id int64,field,path string) (err error) {
	return bad.UpdateArticleImageOrVideo(id,field,path)
}