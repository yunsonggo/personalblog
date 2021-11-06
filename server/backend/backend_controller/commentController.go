package backend_controller

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_service"
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/backend_param"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var commentServer = backend_service.NewBackendCommentServer()
// 获取所有评论数据
func PostCommentAll(ctx *gin.Context) {
	res,err := commentServer.QueryCommentAll()
	if err != nil {
		common.Failed(ctx,"获取评论数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 根据ID 删除评论
func PostRemoveComment(ctx *gin.Context) {
	var bcParam backend_param.BackendCommentParam
	err := ctx.ShouldBindBodyWith(&bcParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取评论参数失败")
		return
	}
	if bcParam.CommentId > 0 {
		err = commentServer.RemoveCommentChildById(bcParam.Id)
	} else {
		err = commentServer.RemoveCommentByID(bcParam.Id)
	}
	if err != nil {
		common.Failed(ctx,"删除评论失败"+err.Error())
		return
	}
	common.Success(ctx,"删除成功")
	return
}