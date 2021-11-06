package backend_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"github.com/gin-gonic/gin"
)

// 获取图形验证码
func GetImgCaptcha(ctx *gin.Context) {
	captchaBody,err := common.GenerateCaptchaCode()
	if err != nil {
		common.Failed(ctx,"生成验证码失败")
		return
	}
	common.Success(ctx,captchaBody)
	return
}