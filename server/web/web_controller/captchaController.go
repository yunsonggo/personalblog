package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/web_param"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func PostEmailCaptcha(ctx *gin.Context) {
	var email web_param.EmailParam
	err := ctx.ShouldBind(&email)
	if err != nil {
		common.Failed(ctx, "获取注册邮箱失败")
		return
	}
	var toEmail = email.Email
	// 发送邮箱验证码
	theCode,err := common.SendEmailCode(toEmail)
	if err != nil {
		common.Failed(ctx, "发送邮箱验证码失败,请重试")
		return
	}
	// 设置session
	ssErr := common.SetSess(ctx,"email_code_" + toEmail,theCode,180)
	if ssErr != nil {
		common.Failed(ctx, "设置邮箱验证码session错误")
		return
	}
	common.Success(ctx,theCode)
	return
}

type SmsData struct {
	BizId string `json:"bizId"`
	Code string `json:"code"`
}

func PostPhoneCaptcha(ctx *gin.Context) {
	var phoneParam web_param.PhoneParam
	err := ctx.ShouldBind(&phoneParam)
	if err != nil {
		common.Failed(ctx,"获取注册手机号失败")
		return
	}
	smsServer := common.NewSmsServer()
	bizId,code ,err := smsServer.ImitateSms(phoneParam.Phone,phoneParam.TplId,phoneParam.Key)
	if err != nil {
		common.Failed(ctx,"发送短信验证码失败")
		return
	}
	var smsData SmsData
	smsData.BizId = bizId
	smsData.Code = code
	jsonSms,_ := json.Marshal(smsData)
	smsStr := string(jsonSms)
	ssErr := common.SetSess(ctx,"phone_code_" + phoneParam.Phone,smsStr,180)
	if ssErr != nil {
		common.Failed(ctx, "设置验证码session错误")
		return
	}
	common.Success(ctx,smsData)
	return
}

func GetImgCaptcha(ctx *gin.Context) {
	captchaBody,err := common.GenerateCaptchaCode()
	if err != nil {
		common.Failed(ctx,"生成验证码失败")
		return
	}
	common.Success(ctx,captchaBody)
	return
}