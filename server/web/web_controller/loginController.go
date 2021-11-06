package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/web_param"
	"2021/yunsongcailu/yunsong_server/tools"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

var cs = web_service.NewConsumerService()
// 邮箱登录
func PostLoginByEmail(ctx *gin.Context) {
	var emailLoginParam web_param.EmailLoginParam
	err := ctx.ShouldBind(&emailLoginParam)
	if err != nil {
		common.Failed(ctx,"获取登录参数失败")
		return
	}
	if !tools.VerifyEmailFormat(emailLoginParam.Email) {
		common.Failed(ctx,"邮箱格式错误")
		return
	}
	if !tools.VerifyPassword(emailLoginParam.Password) {
		common.Failed(ctx,"密码格式错误")
		return
	}
	var code common.VerifyCaptchaBody
	code.Id = emailLoginParam.CaptchaId
	code.VerifyValue = emailLoginParam.Captcha
	if  !common.VerifyCaptchaCode(code) {
		common.Failed(ctx,"验证码错误,点击验证码刷新")
		return
	}
	res,err := cs.GetConsumerByEmail(emailLoginParam.Email)
	if err != nil {
		common.Failed(ctx,err.Error()+"邮箱验证失败")
		return
	}
	if res.Id == 0 {
		common.Failed(ctx,"邮箱未注册")
		return
	}
	loginPassword := tools.EncodeSha256(emailLoginParam.Password)
	if loginPassword != res.Password {
		common.Failed(ctx,"密码错误")
		return
	}
	// 生成token
	token,tokenErr := common.ReleaseToken(res.Id)
	if tokenErr != nil {
		common.Failed(ctx,"生成token失败,请重新登录")
		return
	}
	// 设置session
	idStr := strconv.FormatInt(res.Id,10)
	sessErr := common.SetSess(ctx,"consumer_"+idStr,idStr,0)
	if sessErr != nil {
		common.Failed(ctx,"设置session失败,请重新登录")
		return
	}
	res.Password = token
	common.Success(ctx,res)
}
// 手机登录
func PostLoginByPhone(ctx *gin.Context) {
	var phoneLoginParam web_param.PhoneLoginParam
	err := ctx.ShouldBind(&phoneLoginParam)
	if err != nil {
		common.Failed(ctx,"获取登录参数失败")
		return
	}
	var smsData SmsData
	jsonSms := common.GetSess(ctx,"phone_code_"+phoneLoginParam.Phone)
	strSms := jsonSms.(string)
	err = json.Unmarshal([]byte(strSms),&smsData)
	if err != nil {
		common.Failed(ctx,"解析session错误")
		return
	}
	if smsData.BizId != phoneLoginParam.BizId {
		common.Failed(ctx,"验证码特征异常,请重新获取验证码")
		return
	}
	if smsData.Code != phoneLoginParam.Code {
		common.Failed(ctx,"验证码错误")
		return
	}
	// 查询手机
	res,err := cs.GetConsumerByPhone(phoneLoginParam.Phone)
	if err != nil {
		common.Failed(ctx,err.Error()+"手机验证失败")
		return
	}
	if res.Id == 0 {
		common.Failed(ctx,"手机号未注册")
		return
	}
	// 生成token
	token,tokenErr := common.ReleaseToken(res.Id)
	if tokenErr != nil {
		common.Failed(ctx,"生成token失败,请重新登录")
		return
	}
	// 设置session
	idStr := strconv.FormatInt(res.Id,10)
	sessErr := common.SetSess(ctx,"consumer_"+idStr,idStr,0)
	if sessErr != nil {
		common.Failed(ctx,"设置session失败,请重新登录")
		return
	}
	res.Password = token
	common.Success(ctx,res)
}