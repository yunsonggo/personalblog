package backend_service

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_dao"
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
	"2021/yunsongcailu/yunsong_server/web/web_dao"
)

type BackendCommentServer interface {
	// 获取所有评论
	QueryCommentAll() (commentList []backend_model.BackendComment,err error)
	// 根据ID 删除评论
	RemoveCommentByID(id int64) (err error)
	// 根据ID 删除二级评论
	RemoveCommentChildById(id int64) (err error)
}

type backendCommentServer struct {}

func NewBackendCommentServer() BackendCommentServer {
	return &backendCommentServer{}
}

var bCDao = backend_dao.NewBackendCommentDao()
var commentChildDao = web_dao.NewCommentChildDao()
var articleDao = web_dao.NewArticleDao()
// 获取所有评论
func (bcs *backendCommentServer) QueryCommentAll() (commentList []backend_model.BackendComment,err error) {
	commentArr,err := bCDao.FindCommentAll()
	if err != nil {
		return
	}
	for _,commentPer := range commentArr {
		var comment backend_model.BackendComment
		comment.Id = commentPer.Id
		comment.ArticleId = commentPer.ArticleId
		comment.AuthorId = commentPer.AuthorId
		comment.AuthorNickname = commentPer.AuthorNickname
		comment.AuthorIcon = commentPer.AuthorIcon
		comment.Content = commentPer.Content
		comment.CreateTime = commentPer.CreateTime
		children ,childErr := commentChildDao.QueryCommentChildByCommentId(commentPer.Id)
		if childErr != nil {
			continue
		}
		if len(children) > 0 {
			comment.HasChild = true
			for _,child := range children {
				var commentChild backend_model.BackendCommentChild
				commentChild.Id = child.Id
				commentChild.CommentId = child.CommentId
				commentChild.AuthorId = child.AuthorId
				commentChild.AuthorNickname = child.AuthorNickname
				commentChild.AuthorIcon = child.AuthorIcon
				commentChild.Content = child.Content
				commentChild.CreateTime = child.CreateTime
				comment.CommentChildren = append(comment.CommentChildren, commentChild)
			}
		} else {
			comment.HasChild = false
		}
		article,articleErr := articleDao.QueryArticleById(commentPer.ArticleId)
		if articleErr != nil {
			continue
		}
		comment.CommentArticle = article
		commentList = append(commentList, comment)
	}
	return
}
// 根据ID 删除评论
func (bcs *backendCommentServer) RemoveCommentByID(id int64) (err error) {
	return bCDao.DeleteCommentByID(id)
}
// 根据ID 删除二级评论
func (bcs *backendCommentServer) RemoveCommentChildById(id int64) (err error) {
	return bCDao.DeleteCommentChildById(id)
}