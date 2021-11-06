package web_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type CommentDao interface {
	// 插入一条评论
	InsertCommentOne(comment *web_model.CommentModel) (id int64,err error)
	// 根据文章ID 获取评论
	QueryCommentByArticleId(articleId int64,count,start int) (commentData []web_model.CommentModel,err error)
}

type commentDao struct {}

func NewCommentDao() CommentDao {
	return &commentDao{}
}

// 插入一条评论
func (cmtD *commentDao) InsertCommentOne(comment *web_model.CommentModel) (id int64,err error) {
	return dial.DB.InsertOne(comment)
}
// 根据文章ID 获取评论
func (cmtD *commentDao) QueryCommentByArticleId(articleId int64,count,start int) (commentData []web_model.CommentModel,err error) {
	err = dial.DB.Where("article_id = ?",articleId).OrderBy("create_time desc").Limit(count,start).Find(&commentData)
	return
}