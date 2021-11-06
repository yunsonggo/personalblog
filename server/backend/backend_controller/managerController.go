package backend_controller

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
	"2021/yunsongcailu/yunsong_server/backend/backend_service"
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/backend_param"
	"2021/yunsongcailu/yunsong_server/tools"
	"github.com/gin-gonic/gin"
	"strconv"
)
var bms = backend_service.NewManagerServer()
// 用户注册
func PostRegisterManager(ctx *gin.Context) {
	var managerParam backend_param.ManagerParam
	err := ctx.ShouldBind(&managerParam)
	if err != nil {
		common.Failed(ctx,"获取注册数据失败")
		return
	}
	// 验证验证码
	var code common.VerifyCaptchaBody
	code.Id = managerParam.CaptchaId
	code.VerifyValue = managerParam.Captcha
	if  !common.VerifyCaptchaCode(code) {
		common.Failed(ctx,"验证码错误,点击验证码刷新")
		return
	}
	//查询是否注册
	res,err := bms.GetManagerByPassword(managerParam.Name)
	if err != nil {
		common.Failed(ctx,"验证账号是否注册出现错误")
		return
	}
	if res.Id > 0 {
		common.Failed(ctx,"账号已经存在")
		return
	}
	var manager backend_model.ManagerModel
	manager.ManagerName = managerParam.Name
	password := tools.EncodeSha256(managerParam.Password)
	manager.ManagerPassword = password
	manager.ManagerPower = 0
	err = bms.AddManager(manager)
	if err != nil {
		common.Failed(ctx,"注册用户失败")
		return
	}
	common.Success(ctx,"注册成功")
	return
}
// 用户登录
func PostLoginManager(ctx *gin.Context) {
	var managerParam backend_param.ManagerParam
	err := ctx.ShouldBind(&managerParam)
	if err != nil {
		common.Failed(ctx,"获取登录参数失败")
		return
	}
	// 验证验证码
	var code common.VerifyCaptchaBody
	code.Id = managerParam.CaptchaId
	code.VerifyValue = managerParam.Captcha
	if  !common.VerifyCaptchaCode(code) {
		common.Failed(ctx,"验证码错误,点击验证码刷新")
		return
	}
	res,err := bms.GetManagerByPassword(managerParam.Name)
	if err != nil || res.Id == 0{
		common.Failed(ctx,"验证账号失败")
		return
	}
	password := tools.EncodeSha256(managerParam.Password)
	if res.ManagerPassword != password {
		common.Failed(ctx,"密码错误")
		return
	}
	token,err := common.ReleaseToken(res.Id)
	if err != nil {
		common.Failed(ctx,"发放token失败")
		return
	}
	res.ManagerPassword = token
	idStr := strconv.FormatInt(res.Id,10)
	err = common.SetSess(ctx,"manager_" + idStr,idStr,0)
	if err != nil {
		common.Failed(ctx,"设置session失败")
		return
	}
	common.Success(ctx,res)
}