package web_service

import (
	"2021/yunsongcailu/yunsong_server/web/web_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type CommentServer interface {
	// 插入一条评论
	AddCommentOne(comment *web_model.CommentModel) (id int64,err error)
	// 根据文章ID 获取评论
	GetCommentByArticleId(articleId int64,count,start int) (commentData []ArticleCommentList,err error)
}

type commentServer struct {}

func NewCommentServer() CommentServer {
	return &commentServer{}
}

type ArticleCommentList struct {
	Comment web_model.CommentModel
	CommentChildList []web_model.CommentChildModel
}

var cmtD = web_dao.NewCommentDao()
var as = NewArticleServer()
// 插入一条评论
func (cmtS *commentServer) AddCommentOne(comment *web_model.CommentModel) (id int64,err error) {
	articleId := comment.ArticleId
	id,err = cmtD.InsertCommentOne(comment)
	if err != nil {
		return
	}
	article,articleErr := as.GetArticleById(articleId)
	if articleErr != nil {
		return
	}
	newArticle := new(web_model.ArticleModel)
	newArticle.CommentNum = article.CommentNum + 1
	err = as.EditArticleOnlyValueById(articleId,"comment_num",newArticle)
	return
}
// 根据文章ID 获取评论
func (cmtS *commentServer) GetCommentByArticleId(articleId int64,count,start int) (commentData []ArticleCommentList,err error){
	commentRes,commentResErr := cmtD.QueryCommentByArticleId(articleId,count,start)
	if commentResErr != nil {
		err = commentResErr
		return
	}
	for _,commentItem := range commentRes {
		var articleCommentList ArticleCommentList
		articleCommentList.Comment = commentItem
		commentChildRes,commentChildResErr := cmtDC.QueryCommentChildByCommentId(commentItem.Id)
		if commentChildResErr != nil {
			err = commentChildResErr
			continue
		}
		articleCommentList.CommentChildList = commentChildRes
		commentData = append(commentData, articleCommentList)
	}
	return
}