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
	"strings"
	"time"
)

var backendArticleServer = backend_service.NewBackendArticleServer()
// 获取所有文章
func PostArticleAll(ctx *gin.Context) {
	res,err := backendArticleServer.FindArticleAll()
	if err != nil {
		common.Failed(ctx,err.Error())
		return
	}
	common.Success(ctx,res)
}
// 根据ID 删除一篇文章
func PostArticleRemove(ctx *gin.Context) {
	var removeParam backend_param.BackendArticleRemoveParam
	err := ctx.ShouldBindBodyWith(&removeParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,err.Error())
		return
	}
	id,_ := strconv.ParseInt(removeParam.IdStrArr,10,64)
	err = backendArticleServer.RemoveArticleById(id)
	if err != nil {
		common.Failed(ctx,err.Error())
		return
	}
	common.Success(ctx,"删除成功")
	return
}
// 批量删除文章
func PostRemoveArticleArr(ctx *gin.Context) {
	var removeParam backend_param.BackendArticleRemoveParam
	err := ctx.ShouldBindBodyWith(&removeParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,err.Error())
		return
	}
	idStrArr := strings.Split(removeParam.IdStrArr,",")
	var idIntArr []int64
	for _,idStr := range idStrArr {
		id,_ := strconv.ParseInt(idStr,10,64)
		if id == 0 {
			continue
		}
		idIntArr = append(idIntArr, id)
	}
	err = backendArticleServer.RemoveArticleArrById(idIntArr)
	if err != nil {
		common.Failed(ctx,err.Error())
		return
	}
	common.Success(ctx," 删除成功")
	return
}
// 上传文章banner
func PostArticleUploadBanner(ctx *gin.Context) {
	bannerParamData,_ := ctx.Get("bannerParam")
	if bannerParamData == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	bannerParam := bannerParamData.(backend_param.ArticleBannerUploadParam)
	bannerFileHeader := bannerParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(bannerFileHeader.Filename)
	fileName := "backendUpload/article/banner/banner" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(bannerFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存Banner失败")
		return
	}
	// 如果存在旧文件
	oldBanner := bannerParam.OldBanner
	if oldBanner != "" {
		oldBannerPath := "public/" + oldBanner
		_ = os.Remove(oldBannerPath)
	}
	// 如果存在文章ID 则更新数据库 如果不存在则不更新
	if bannerParam.ArticleId > 0 {
		// 更新数据
		_ = backendArticleServer.EditArticleImageOrVideo(bannerParam.ArticleId,"banner",fileName)
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}
// 上传文章封面cover
func PostArticleUploadCover(ctx *gin.Context) {
	coverParamData,_ := ctx.Get("coverParam")
	if coverParamData == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	coverParam := coverParamData.(backend_param.ArticleCoverUploadParam)
	coverFileHeader := coverParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(coverFileHeader.Filename)
	fileName := "backendUpload/article/cover/cover" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(coverFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存cover失败")
		return
	}
	// 如果存在旧文件
	oldCover := coverParam.OldCover
	if oldCover != "" {
		oldCoverPath := "public/" + oldCover
		_ = os.Remove(oldCoverPath)
	}
	// 如果存在文章ID 则更新数据库 如果不存在则不更新
	if coverParam.ArticleId > 0 {
		// 更新数据
		_ = backendArticleServer.EditArticleImageOrVideo(coverParam.ArticleId,"cover",fileName)
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}
// 富文本编辑器上传图片
func PostEditorUploadImg(ctx *gin.Context) {
	_,fileHeader,err := ctx.Request.FormFile("file")
	if err != nil {
		common.Failed(ctx,err.Error())
		return
	}
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(fileHeader.Filename)
	fileName := "backendUpload/article/editor/editor" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err = ctx.SaveUploadedFile(fileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存cover失败")
		return
	}
	common.Success(ctx,fileName)
	return
}
// 上传文章视频
func PostArticleUploadVideo(ctx *gin.Context) {
	videoParamData,_ := ctx.Get("videoParam")
	if videoParamData == nil {
		common.Failed(ctx,"获取上传参数失败")
		return
	}
	videoParam := videoParamData.(backend_param.ArticleVideoUploadParam)
	videoFileHeader := videoParam.File
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(videoFileHeader.Filename)
	fileName := "backendUpload/article/video/video" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(videoFileHeader,filePath)
	if err != nil {
		common.Failed(ctx,"保存cover失败")
		return
	}
	// 如果存在旧文件
	oldVideo := videoParam.OldVideo
	if oldVideo != "" {
		oldVideoPath := "public/" + oldVideo
		_ = os.Remove(oldVideoPath)
	}
	// 如果存在文章ID 则更新数据库 如果不存在则不更新
	if videoParam.ArticleId > 0 {
		// 更新数据
		_ = backendArticleServer.EditArticleImageOrVideo(videoParam.ArticleId,"video",fileName)
		common.Success(ctx,fileName)
		return
	} else {
		common.Success(ctx,fileName)
		return
	}
}
// 添加新文章
func PostArticleAdd(ctx *gin.Context) {
	var articleAddParam backend_param.ArticleAddParam
	err := ctx.ShouldBindBodyWith(&articleAddParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取文章参数失败")
		return
	}
	article := articleAddParam.ArticleInfo
	err = backendArticleServer.AddArticle(article)
	if err != nil {
		common.Failed(ctx,"添加文章失败")
		return
	}
	common.Success(ctx,"添加文章成功")
	return
}
// 根据ID 获取文章数据
func PostArticleInfo(ctx *gin.Context) {
	var articleEditParam backend_param.ArticleEditParam
	err := ctx.ShouldBindBodyWith(&articleEditParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取文章ID失败")
		return
	}
	articleId := articleEditParam.ArticleId
	res,err := backendArticleServer.GetArticleById(articleId)
	if err != nil {
		common.Failed(ctx,"获取文章数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 修改文章
func PostArticleEdit(ctx *gin.Context) {
	var articleEditParam backend_param.ArticleEditParam
	err := ctx.ShouldBindBodyWith(&articleEditParam,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取文章数据失败")
		return
	}
	articleInfo := articleEditParam.ArticleInfo
	err = backendArticleServer.EditArticleById(articleInfo)
	if err != nil {
		common.Failed(ctx,"更新文章失败")
		return
	}
	common.Success(ctx,"更新文章成功")
	return
}
