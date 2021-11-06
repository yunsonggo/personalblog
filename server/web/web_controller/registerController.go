package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/web_param"
	"2021/yunsongcailu/yunsong_server/tools"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

var consumerServer = web_service.NewConsumerService()

func PostRegisterByEmail(ctx *gin.Context) {
	var registerParam web_param.RegisterEmailParam
	err := ctx.ShouldBind(&registerParam)
	if err != nil {
		common.Failed(ctx,"请输入必填项")
		return
	}
	toEmail := registerParam.Email
	sessCode := common.GetSess(ctx,"email_code_"+toEmail)
	if sessCode != registerParam.Code {
		common.Failed(ctx,"验证码错误,请重新获取验证码!")
		return
	}
	if !tools.VerifyEmailFormat(registerParam.Email) {
		common.Failed(ctx,"邮箱格式错误")
		return
	}
	if !tools.VerifyPassword(registerParam.Password) {
		common.Failed(ctx,"密码格式错误")
		return
	}
	if registerParam.Password != registerParam.CheckPass {
		common.Failed(ctx,"两次密码不一致,请检查输入")
		return
	}
	//if !tools.VerifyChineseWord(registerParam.Name) {
	//	common.Failed(ctx,"姓名不是中文字符")
	//	return
	//}
	// 验证是否已经注册
	res ,err := consumerServer.GetConsumerByEmail(registerParam.Email)
	if err != nil {
		common.Failed(ctx,"验证邮箱是否注册出现错误,请重新注册")
		return
	}
	if res.Id != 0 {
		common.Failed(ctx,"邮箱:"+toEmail+"已经注册了,请登录")
		return
	}
	newConsumer,err := consumerServer.AddConsumer(registerParam.Name,registerParam.Email,registerParam.Password,"")
	if err != nil || newConsumer.Id == 0 {
		fmt.Printf("err = %v\n",err)
		common.Failed(ctx,"注册新用户失败")
		return
	}
	// 生成token
	token,err := common.ReleaseToken(newConsumer.Id)
	if err != nil {
		common.Failed(ctx,"用户注册成功,发放token失败,请重新登录!")
		return
	}
	// 设置session
	idStr := strconv.FormatInt(newConsumer.Id,10)
	err = common.SetSess(ctx,"consumer_"+idStr,idStr,3600*24*7)
	if err != nil {
		common.Failed(ctx,"用户注册成功,设置session失败,请重新登录!")
		return
	}
	newConsumer.Password = token
	common.Success(ctx,newConsumer)
	return
}

func PostRegisterByPhone(ctx *gin.Context) {
	var registerParam web_param.RegisterPhoneParam
	err := ctx.ShouldBind(&registerParam)
	if err != nil {
		common.Failed(ctx,"请输入必填项")
		return
	}
	var smsData SmsData
	jsonSms := common.GetSess(ctx,"phone_code_"+registerParam.Phone)
	strSms := jsonSms.(string)
	err = json.Unmarshal([]byte(strSms),&smsData)
	if err != nil {
		common.Failed(ctx,"解析session错误")
		return
	}
	if smsData.BizId != registerParam.BizId {
		common.Failed(ctx,"验证码特征异常,请重新获取验证码")
		return
	}
	if smsData.Code != registerParam.Code {
		common.Failed(ctx,"验证码错误")
		return
	}
	if !tools.VerifyMobileFormat(registerParam.Phone) {
		common.Failed(ctx,"手机号码格式错误")
		return
	}
	// 验证手机是否注册
	res,err := consumerServer.GetConsumerByPhone(registerParam.Phone)
	if err != nil {
		common.Failed(ctx,"验证手机号码是否已经注册错误,请重新注册")
		return
	}
	if res.Id != 0 {
		common.Failed(ctx,"手机号码"+registerParam.Phone+"已经注册,请登录")
		return
	}
	newConsumer,err := consumerServer.AddConsumer(registerParam.Name,"","",registerParam.Phone)
	if err != nil || newConsumer.Id == 0 {
		fmt.Printf("err = %v\n",err)
		common.Failed(ctx,"注册新用户失败")
		return
	}
	// 生成token
	token,err := common.ReleaseToken(newConsumer.Id)
	if err != nil {
		common.Failed(ctx,"用户注册成功,发放token失败,请重新登录!")
		return
	}
	// 设置session
	idStr := strconv.FormatInt(newConsumer.Id,10)
	err = common.SetSess(ctx,"consumer_"+idStr,idStr,3600*24*7)
	if err != nil {
		common.Failed(ctx,"用户注册成功,设置session失败,请重新登录!")
		return
	}
	newConsumer.Password = token
	common.Success(ctx,newConsumer)
	return
}