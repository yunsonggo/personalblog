package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/web_param"
	"2021/yunsongcailu/yunsong_server/web/web_model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// 上传更新头像
func AuthPostPersonalUploadAvatar(ctx *gin.Context) {
	//_,avatarFileHeader,err := ctx.Request.FormFile("file")
	avatarParamData,_ := ctx.Get("avatarParam")
	if avatarParamData == nil {
		common.Failed(ctx,"获取上传文件失败")
		return
	}
	avatarParam := avatarParamData.(web_param.PersonalAvatarParams)
	avatarFileHeader := avatarParam.File
	// 随机数
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(avatarFileHeader.Filename)
	fileName := "webUpload/consumer/avatar/avatar" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(avatarFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存文件失败")
		return
	}
	// 如果存在旧文件
	oldAvatar := avatarParam.OldAvatar
	if oldAvatar != ""{
		oldFilePath := "public/" + oldAvatar
		_ = os.Remove(oldFilePath)
	}
	// 更新数据库
	webUser,isSet := ctx.Get("webUser")
	var userInfo web_model.Consumers
	if isSet {
		userInfo = webUser.(web_model.Consumers)
	}
	err = cs.EditConsumerIconById(userInfo.Id,fileName)
	if err != nil {
		common.Failed(ctx,"更新头像入库失败")
		return
	}
	common.Success(ctx,fileName)
}

// 修改个人信息
func AuthPostPersonalEdit(ctx *gin.Context) {
	var consumerEditParam web_param.PersonalEditParam
	err := ctx.ShouldBindBodyWith(&consumerEditParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取用户数据失败")
		return
	}
	consumerInfo := consumerEditParam.Consumer
	var consumer web_model.Consumers
	consumer.Id = consumerInfo.Id
	if consumerEditParam.IsEmail == "false" {
		consumer.Email = consumerInfo.Email
	}
	if consumerEditParam.IsPhone == "false" {
		consumer.Phone = consumerInfo.Phone
	}
	consumer.Gender = consumerInfo.Gender
	consumer.NickName = consumerInfo.Nickname
	consumer.Desc = consumerInfo.Desc
	err = cs.EditConsumerInfoById(consumer)
	if err != nil {
		common.Failed(ctx,"更新用户信息失败")
		return
	}
	common.Success(ctx,"OK")
	return
}