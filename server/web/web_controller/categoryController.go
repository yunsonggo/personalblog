package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"github.com/gin-gonic/gin"
)

var cgs = web_service.NewCategoryServer()
// 获取所有激活类别
func PostFindCategory(ctx *gin.Context) {
	res,err := cgs.FindCategory()
	if err != nil {
		common.Failed(ctx,"获取类别数据失败")
		return
	}
	common.Success(ctx,res)
}