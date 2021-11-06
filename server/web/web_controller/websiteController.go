package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"github.com/gin-gonic/gin"
)
var ws = web_service.NewWebsiteServer()
// 获取网站信息
func PostWebsiteInfo(ctx *gin.Context) {
	res,err := ws.GetWebsiteInfo()
	if err != nil {
		common.Failed(ctx,"获取网站信息失败")
		return
	}
	common.Success(ctx,res)
}