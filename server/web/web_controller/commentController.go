package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/web_param"
	"2021/yunsongcailu/yunsong_server/web/web_model"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"github.com/gin-gonic/gin"
)
var cmtS = web_service.NewCommentServer()
var cmtSC = web_service.NewCommentChildServer()
// 根据文章ID 获取评论
func PostCommentByArticleId(ctx *gin.Context) {
	var commentParam web_param.CommentParam
	err := ctx.ShouldBind(&commentParam)
	if err != nil {
		common.Failed(ctx,"获取文章参数失败")
		return
	}
	res,err := cmtS.GetCommentByArticleId(commentParam.ArticleId,commentParam.Count,commentParam.Start)
	if err != nil {
		common.Failed(ctx,"获取文章评论失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 根据文章ID 添加评论
func PostAddComment(ctx *gin.Context) {
	var commentParam web_param.CommentParam
	err := ctx.ShouldBind(&commentParam)
	if err != nil {
		common.Failed(ctx,"获取评论参数失败")
		return
	}
	var comment web_model.CommentModel
	comment.ArticleId = commentParam.ArticleId
	comment.AuthorId = commentParam.AuthorId
	comment.Content = commentParam.Content
	comment.AuthorNickname = commentParam.AuthorNickname
	comment.AuthorIcon = commentParam.AuthorIcon
	_,err = cmtS.AddCommentOne(&comment)
	if err != nil {
		common.Failed(ctx,"添加评论失败")
		return
	}
	common.Success(ctx,comment)
	return
}
// 添加二级评论
func PostAddCommentChild(ctx *gin.Context) {
	var commentChildParam web_param.CommentChildParam
	err := ctx.ShouldBind(&commentChildParam)
	if err != nil {
		common.Failed(ctx,"获取二级评论参数失败")
		return
	}
	var commentChild web_model.CommentChildModel
	commentChild.CommentId = commentChildParam.CommentId
	commentChild.AuthorId = commentChildParam.AuthorId
	commentChild.Content = commentChildParam.Content
	commentChild.AuthorNickname = commentChildParam.AuthorNickname
	commentChild.AuthorIcon = commentChildParam.AuthorIcon
	_,err = cmtSC.AddCommentChildOne(&commentChild)
	if err != nil {
		common.Failed(ctx,"添加二级评论失败")
		return
	}
	common.Success(ctx,commentChild)
	return
}