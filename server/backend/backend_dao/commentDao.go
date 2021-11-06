package backend_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type BackendCommentDao interface {
	// 获取所有评论
	FindCommentAll() (comments []web_model.CommentModel,err error)
	// 根据ID 删除评论
	DeleteCommentByID(id int64) (err error)
	// 根据ID 删除二级评论
	DeleteCommentChildById(id int64) (err error)
}

type backendCommentDao struct {}

func NewBackendCommentDao() BackendCommentDao {
	return &backendCommentDao{}
}

// 获取所有评论
func (bcd *backendCommentDao) FindCommentAll() (comments []web_model.CommentModel,err error) {
	err = dial.DB.Desc("create_time").Find(&comments)
	return
}
// 根据ID 删除评论
func (bcd *backendCommentDao) DeleteCommentByID(id int64) (err error) {
	comment := new(web_model.CommentModel)
	_,err = dial.DB.Where("id = ?",id).Delete(comment)
	return
}
// 根据ID 删除二级评论
func (bcd *backendCommentDao) DeleteCommentChildById(id int64) (err error) {
	commentChild := new(web_model.CommentChildModel)
	_,err = dial.DB.Where("id = ?",id).Delete(commentChild)
	return
}