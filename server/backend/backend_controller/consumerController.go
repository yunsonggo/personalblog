package backend_controller

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_service"
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/backend_param"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var backendConsumerServer = backend_service.NewConsumerServer()

// 用户列表
func PostConsumerList(ctx *gin.Context) {
	res,err := backendConsumerServer.FindConsumers()
	if err != nil {
		common.Failed(ctx,"获取用户列表失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 修改用户状态 1 删除 0 激活
func PostConsumerState(ctx *gin.Context) {
	var backendConsumerParam backend_param.BackendConsumerParam
	err := ctx.ShouldBindBodyWith(&backendConsumerParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取用户参数失败")
		return
	}
	err = backendConsumerServer.EditConsumerState(backendConsumerParam.Id,backendConsumerParam.State)
	if err != nil {
		common.Failed(ctx,"更新用户状态失败")
		return
	}
	common.Success(ctx,"成功")
	return
}