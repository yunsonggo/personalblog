package web_service

import (
	"2021/yunsongcailu/yunsong_server/web/web_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type CommentChildServer interface {
	// 插入一条子评论
	AddCommentChildOne(commentChild *web_model.CommentChildModel) (id int64,err error)
	// 根据评论ID 获取子评论
	GetCommentChildByCommentId(commentId int64) (commentChildData []web_model.CommentChildModel,err error)
}

type commentChildServer struct {}

func NewCommentChildServer() CommentChildServer {
	return &commentChildServer{}
}

var cmtDC = web_dao.NewCommentChildDao()
// 插入一条子评论
func (cmtCS *commentChildServer) AddCommentChildOne(commentChild *web_model.CommentChildModel) (id int64,err error) {
	return cmtDC.InsertCommentChildOne(commentChild)
}
// 根据评论ID 获取子评论
func (cmtCS *commentChildServer) GetCommentChildByCommentId(commentId int64) (commentChildData []web_model.CommentChildModel,err error) {
	return cmtDC.QueryCommentChildByCommentId(commentId)
}