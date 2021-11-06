package web_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type CommentChildDao interface {
	// 插入一条子评论
	InsertCommentChildOne(commentChild *web_model.CommentChildModel) (id int64,err error)
	// 根据评论ID 获取子评论
	QueryCommentChildByCommentId(commentId int64) (commentChildData []web_model.CommentChildModel,err error)
}

type commentChildDao struct {}

func NewCommentChildDao() CommentChildDao {
	return &commentChildDao{}
}
// 插入一条子评论
func (cmtDC *commentChildDao) InsertCommentChildOne(commentChild *web_model.CommentChildModel) (id int64,err error) {
	return dial.DB.InsertOne(commentChild)
}
// 根据评论ID 获取子评论
func (cmtDC *commentChildDao) QueryCommentChildByCommentId(commentId int64) (commentChildData []web_model.CommentChildModel,err error) {
	err = dial.DB.Where("comment_id = ?",commentId).Find(&commentChildData)
	return
}