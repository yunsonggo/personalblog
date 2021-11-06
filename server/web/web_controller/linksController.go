package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"github.com/gin-gonic/gin"
)

var ls = web_service.NewLinksServer()
// 获取所有链接
func PostLinks(ctx *gin.Context) {
	res,err := ls.FindLinks()
	if err != nil {
		common.Failed(ctx,"获取友情链接失败")
		return
	}
	common.Success(ctx,res)
}
