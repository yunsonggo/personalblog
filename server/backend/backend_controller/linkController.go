package backend_controller

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_service"
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/backend_param"
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

var bls = backend_service.NewBackendLinkServer()

// 获取所有链接
func PostLinkAll(ctx *gin.Context) {
	res,err := bls.GetLinkAll()
	if err != nil {
		common.Failed(ctx,"获取链接数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 修改链接
func PostLinkEdit(ctx *gin.Context) {
	var linkTextParam backend_param.LinkTextParam
	err := ctx.ShouldBindBodyWith(&linkTextParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取链接参数失败")
		return
	}
	var newLink = web_model.LinksModel{
		Id: linkTextParam.Id,
		LinkIcon: linkTextParam.Icon,
		LinkTitle: linkTextParam.Title,
		LinkUrl: linkTextParam.LinkUrl,
		Sort: linkTextParam.Sort,
	}
	err = bls.EditLinkById(newLink)
	if err != nil {
		common.Failed(ctx,"更新链接失败")
		return
	}
	common.Success(ctx,"OK")
	return
}
// 更新链接图片
func PostLinkIconEdit(ctx *gin.Context) {
	linkParamData,_ := ctx.Get("linkIconParam")
	if linkParamData == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	linkParam := linkParamData.(backend_param.LinkIconUploadParam)
	iconFileHeader := linkParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(iconFileHeader.Filename)
	fileName := "webUpload/links/link" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(iconFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存链接图片失败")
		return
	}
	// 如果存在旧文件
	oldIcon:= linkParam.OldIcon
	if oldIcon != "" {
		oldIconPath := "public/" + oldIcon
		_ = os.Remove(oldIconPath)
	}
	// 如果存在D 则更新数据库 如果不存在则不更新
	if linkParam.LinkId > 0 {
		// 更新数据
		err = bls.EditLinkIcon(linkParam.LinkId,fileName)
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
// 根据ID 删除链接
func PostLinkRemoveById(ctx *gin.Context) {
	var linkParam backend_param.LinkTextParam
	err := ctx.ShouldBindBodyWith(&linkParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取链接参数失败")
		return
	}
	err = bls.RemoveLinkById(linkParam.Id)
	if err != nil {
		common.Failed(ctx,"删除链接失败")
		return
	}
	common.Success(ctx,"删除成功")
	return
}
// 添加链接
func PostLinkInsertOne(ctx *gin.Context) {
	var linkParam backend_param.LinkTextParam
	err := ctx.ShouldBindBodyWith(&linkParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取链接参数失败")
		return
	}
	var link web_model.LinksModel
	link.Id = linkParam.Id
	link.LinkIcon = linkParam.Icon
	link.LinkUrl = linkParam.LinkUrl
	link.LinkTitle = linkParam.Title
	link.Sort = linkParam.Sort
	err = bls.AddLink(link)
	if err != nil {
		common.Failed(ctx,"插入数据库失败")
		return
	}
	common.Success(ctx,"添加链接成功")
	return
}