package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"github.com/gin-gonic/gin"
)

var ms = web_service.NewMenuServer()

// 获取所有激活菜单
func PostMenu(ctx *gin.Context) {
	res,err := ms.FindMenu()
	if err != nil {
		common.Failed(ctx,"获取菜单失败")
		return
	}
	common.Success(ctx,res)
}