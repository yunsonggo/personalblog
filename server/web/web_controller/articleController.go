package web_controller

import (
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/web_param"
	"2021/yunsongcailu/yunsong_server/tools"
	"2021/yunsongcailu/yunsong_server/web/web_model"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)
var as = web_service.NewArticleServer()
//获取第一大类文章
func PostArticleIndex(ctx *gin.Context) {
	var articleParam web_param.ArticleParam
	err := ctx.ShouldBind(&articleParam)
	if err != nil {
		common.Failed(ctx,"获取文章参数失败")
		return
	}
	res ,err := as.FindIndexArticle(articleParam.Count,articleParam.Start)
	if err != nil {
		common.Failed(ctx,"获取文章数据失败")
		return
	}
	common.Success(ctx,res)
}
// 获取非第一大类之外 每个大类的首页展示文章
func PostOtherArticleIndex(ctx *gin.Context) {
	var articleParam web_param.ArticleParam
	err := ctx.ShouldBind(&articleParam)
	if err != nil {
		common.Failed(ctx,"获取文章参数失败")
		return
	}
	res ,err := as.FindMenuIndexArticle(articleParam.Count,articleParam.Start)
	if err != nil {
		common.Failed(ctx,"获取文章数据失败")
		return
	}
	common.Success(ctx,res)
}
// 获取指定大类和小类的文章
func PostMenuIdCategoryIdArticle(ctx *gin.Context) {
	var articleParam web_param.ArticleParam
	err := ctx.ShouldBind(&articleParam)
	if err != nil {
		common.Failed(ctx,"获取文章参数失败")
		return
	}
	res,err := as.FindMenuCategoryArticle(articleParam.MenuId,articleParam.CategoryId,articleParam.Count,articleParam.Start)
	if err != nil {
		common.Failed(ctx,"获取文章数据失败")
		return
	}
	common.Success(ctx,res)
}
// 获取指定大类无小类文章
func PostMenuArticle(ctx *gin.Context) {
	var articleParam web_param.ArticleParam
	err := ctx.ShouldBind(&articleParam)
	if err != nil {
		common.Failed(ctx,"获取文章参数失败")
		return
	}
	res,err := as.FindMenuArticle(articleParam.MenuId,articleParam.Count,articleParam.Start)
	if err != nil {
		common.Failed(ctx,"获取文章数据失败")
		return
	}
	common.Success(ctx,res)
}
// 根据ID 获取文章
func PostArticleById(ctx *gin.Context) {
	var articleParam web_param.ArticleParam
	err := ctx.ShouldBind(&articleParam)
	if err != nil {
		common.Failed(ctx,"获取文章参数失败")
		return
	}
	res,err := as.GetArticleById(articleParam.ArticleId)
	if err != nil {
		common.Failed(ctx,"获取文章数据失败")
		return
	}
	common.Success(ctx,res)
}
// 更具ID 更新 赞或者 踩
func PostArticleGoodOrBad(ctx *gin.Context) {
	var articleParam web_param.ArticleParam
	err := ctx.ShouldBind(&articleParam)
	if err != nil {
		common.Failed(ctx,"获取文章参数失败")
		return
	}
	newArticle := new(web_model.ArticleModel)
	if articleParam.JobTag == "good" {
		newArticle.Star = articleParam.Star + 1
		err = as.EditArticleOnlyValueById(articleParam.ArticleId,"star",newArticle)
	} else if articleParam.JobTag == "bad" {
		newArticle.Tread = articleParam.Tread + 1
		err = as.EditArticleOnlyValueById(articleParam.ArticleId,"tread",newArticle)
	} else {
		err = errors.New("非正常指令")
	}
	if err != nil {
		common.Failed(ctx,err.Error())
		return
	}
	common.Success(ctx,"ok")
	return
}
// 根据关键字搜索文章
func PostSearchArticle(ctx *gin.Context) {
	var articleParam web_param.ArticleParam
	err := ctx.ShouldBind(&articleParam)
	if err != nil {
		common.Failed(ctx,"获取搜索关键词失败")
		return
	}
	var keywordArr []string
	var articleList []web_model.ArticleModel
	arr := strings.Fields(articleParam.SearchKeyword)
	for _,item := range arr {
		itemArr := strings.Split(item,",")
		for _,keyword := range itemArr {
			if keyword == "" {
				continue
			}
			keywordArr = append(keywordArr, keyword)
		}
	}
	keywordArr = tools.RemoveStringByMap(keywordArr)
	for _,keyword := range keywordArr {
		fmt.Println(keyword)
		res,err := as.FindArticleByKeyword("title",keyword,articleParam.Count,articleParam.Start)
		if err != nil {
			continue
		}
		articleList = append(articleList, res...)
	}
	common.Success(ctx,articleList)
}