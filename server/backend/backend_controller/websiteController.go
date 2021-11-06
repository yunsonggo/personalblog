package backend_controller

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_service"
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/backend_param"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var bws = backend_service.NewBackendWebsiteServer()

// 获取网站信息
func PostWebsiteInfo(ctx *gin.Context) {
	res,err := bws.GetWebsite()
	if err != nil {
		common.Failed(ctx,"获取网站信息失败")
		return
	}
	common.Success(ctx,res)
}
// 更新网站头像
func PostWebsiteAvatarEdit(ctx *gin.Context) {
	websiteAvatarParam,_ := ctx.Get("websiteAvatarParam")
	if websiteAvatarParam == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	avatarParam := websiteAvatarParam.(backend_param.WebsiteAvatarParam)
	avatarFileHeader := avatarParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(avatarFileHeader.Filename)
	fileName := "webUpload/website/avatar" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(avatarFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存链接图片失败")
		return
	}
	// 如果存在旧文件
	oldAvatar:= avatarParam.OldAvatar
	if oldAvatar != "" {
		oldAvatarPath := "public/" + oldAvatar
		_ = os.Remove(oldAvatarPath)
	}
	// 如果存在D 则更新数据库 如果不存在则不更新
	if avatarParam.Id > 0 {
		// 更新数据
		err = bws.EditAvatar(avatarParam.Id,fileName)
		if err != nil {
			common.Failed(ctx,"上传成功,但更新数据库失败")
			return
		}
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}
func PostWebsiteLogoEdit(ctx *gin.Context) {
	websiteLogoParam,_ := ctx.Get("websiteLogoParam")
	if websiteLogoParam == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	logoParam := websiteLogoParam.(backend_param.WebsiteLogoParam)
	logoFileHeader := logoParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(logoFileHeader.Filename)
	fileName := "webUpload/website/logo" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(logoFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存链接图片失败")
		return
	}
	// 如果存在旧文件
	oldLogo:= logoParam.OldLogo
	if oldLogo != "" {
		oldLogoPath := "public/" + oldLogo
		_ = os.Remove(oldLogoPath)
	}
	// 如果存在D 则更新数据库 如果不存在则不更新
	if logoParam.Id > 0 {
		// 更新数据
		err = bws.EditLogo(logoParam.Id,fileName)
		if err != nil {
			common.Failed(ctx,"上传成功,但更新数据库失败")
			return
		}
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}
func PostWebsiteBannerEdit(ctx *gin.Context) {
	websiteBannerParam,_ := ctx.Get("websiteBannerParam")
	if websiteBannerParam == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	bannerParam := websiteBannerParam.(backend_param.WebsiteBannerParam)
	bannerFileHeader := bannerParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(bannerFileHeader.Filename)
	fileName := "webUpload/website/banner" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(bannerFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存链接图片失败")
		return
	}
	// 如果存在旧文件
	oldBanner:= bannerParam.OldBanner
	if oldBanner != "" {
		oldBannerPath := "public/" + oldBanner
		_ = os.Remove(oldBannerPath)
	}
	// 如果存在D 则更新数据库 如果不存在则不更新
	if bannerParam.Id > 0 {
		// 更新数据
		err = bws.EditBanner(bannerParam.Id,fileName)
		if err != nil {
			common.Failed(ctx,"上传成功,但更新数据库失败")
			return
		}
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}
func PostWebsiteWeboEdit(ctx *gin.Context) {
	websiteWeboParam,_ := ctx.Get("websiteWeboParam")
	if websiteWeboParam == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	weboParam := websiteWeboParam.(backend_param.WebsiteWeboParam)
	weboFileHeader := weboParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(weboFileHeader.Filename)
	fileName := "webUpload/website/webo" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(weboFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存链接图片失败")
		return
	}
	// 如果存在旧文件
	oldWebo:= weboParam.OldWebo
	if oldWebo != "" {
		oldWeboPath := "public/" + oldWebo
		_ = os.Remove(oldWeboPath)
	}
	// 如果存在D 则更新数据库 如果不存在则不更新
	if weboParam.Id > 0 {
		// 更新数据
		err = bws.EditWebo(weboParam.Id,fileName)
		if err != nil {
			common.Failed(ctx,"上传成功,但更新数据库失败")
			return
		}
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}
func PostWebsiteWecharEdit(ctx *gin.Context) {
	websiteWecharParam,_ := ctx.Get("websiteWecharParam")
	if websiteWecharParam == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	wecharParam := websiteWecharParam.(backend_param.WebsiteWecharParam)
	wecharFileHeader := wecharParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(wecharFileHeader.Filename)
	fileName := "webUpload/website/wechar" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(wecharFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存链接图片失败")
		return
	}
	// 如果存在旧文件
	oldWechar:= wecharParam.OldWechar
	if oldWechar != "" {
		oldWecharPath := "public/" + oldWechar
		_ = os.Remove(oldWecharPath)
	}
	// 如果存在D 则更新数据库 如果不存在则不更新
	if wecharParam.Id > 0 {
		// 更新数据
		err = bws.EditWechar(wecharParam.Id,fileName)
		if err != nil {
			common.Failed(ctx,"上传成功,但更新数据库失败")
			return
		}
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}
// 更新网站基本信息
func PostEditWebsiteInfo(ctx *gin.Context) {
	var websiteParam backend_param.WebsiteParam
	err := ctx.ShouldBindBodyWith(&websiteParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取参数失败")
		return
	}
	err = bws.EditWebsiteInfo(websiteParam.Website)
	if err != nil {
		common.Failed(ctx,"更新数据库错误")
		return
	}
	common.Success(ctx,"更新成功")
	return
}
// 更新网站关于我们
func PostEditWebsiteAbout(ctx *gin.Context) {
	var aboutParam backend_param.AboutParam
	err := ctx.ShouldBindBodyWith(&aboutParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取参数失败")
		return
	}
	err = bws.EditWebsiteText("about",aboutParam.About)
	if err != nil {
		common.Failed(ctx,"更新数据库失败")
		return
	}
	common.Success(ctx,"更新成功")
	return
}
// 更新用户协议
func PostEditWebsiteAgreement(ctx *gin.Context) {
	var agreementParam backend_param.AgreementParam
	err := ctx.ShouldBindBodyWith(&agreementParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取参数失败")
		return
	}
	err = bws.EditWebsiteText("agreement",agreementParam.Agreement)
	if err != nil {
		common.Failed(ctx,"更新数据库失败")
		return
	}
	common.Success(ctx,"更新成功")
	return
}
// 更新隐私声明
func PostEditWebsitePrivacy(ctx *gin.Context) {
	var privacyParam backend_param.PrivacyParam
	err := ctx.ShouldBindBodyWith(&privacyParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取参数失败")
		return
	}
	err = bws.EditWebsiteText("privacy",privacyParam.Privacy)
	if err != nil {
		common.Failed(ctx,"更新数据库失败")
		return
	}
	common.Success(ctx,"更新成功")
	return
}
// 更新APK文件
func PostWebsiteApkEdit(ctx *gin.Context) {
	websiteApkParam,_ := ctx.Get("websiteApkParam")
	if websiteApkParam == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	apkParam := websiteApkParam.(backend_param.WebsiteApkParam)
	apkFileHeader := apkParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(apkFileHeader.Filename)
	fileName := "webUpload/website/apk" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(apkFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存apk失败")
		return
	}
	// 如果存在旧文件
	oldApk:= apkParam.OldApk
	if oldApk != "" {
		oldApkPath := "public/" + oldApk
		_ = os.Remove(oldApkPath)
	}
	// 如果存在D 则更新数据库 如果不存在则不更新
	if apkParam.Id > 0 {
		// 更新数据
		err = bws.EditWebsiteText("apk",fileName)
		if err != nil {
			common.Failed(ctx,"上传成功,但更新数据库失败")
			return
		}
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}