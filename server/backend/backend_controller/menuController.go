package backend_controller

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_service"
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/param/backend_param"
	"2021/yunsongcailu/yunsong_server/web/web_model"
	"2021/yunsongcailu/yunsong_server/web/web_service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var backendMenuServer = backend_service.NewBackendMenuServer()
var webMenuServer = web_service.NewMenuServer()
var webCategoryServer = web_service.NewCategoryServer()
// 获取后台菜单
func PostBackendMenu(ctx *gin.Context) {
	menus,err := backendMenuServer.GetMenus()
	if err != nil {
		common.Failed(ctx,"获取后台菜单失败")
		return
	}
	common.Success(ctx,menus)
}
// 获取所有前台菜单
func PostWebMenu(ctx *gin.Context) {
	webMenus,err := backendMenuServer.FindWebMenuAll()
	if err != nil {
		common.Failed(ctx,"获取前台菜单失败")
		return
	}
	common.Success(ctx,webMenus)
}
// 修改前台菜单
func PostEditWebMenu(ctx *gin.Context) {
	var webMenuParams backend_param.WebMenuParams
	err := ctx.ShouldBindBodyWith(&webMenuParams,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取菜单参数失败")
		return
	}
	if webMenuParams.ParentId == 0 {
		// 修改主菜单
		var newMenu web_model.MenuModel
		newMenu.Id = webMenuParams.Id
		newMenu.Title = webMenuParams.Title
		newMenu.Sort = webMenuParams.Sort
		newMenu.State = webMenuParams.State
		err = webMenuServer.EditMenu(newMenu)
	} else {
		// 修改类别 二级菜单
		var newCategory web_model.CategoryModel
		newCategory.Id = webMenuParams.Id
		newCategory.Title = webMenuParams.Title
		newCategory.State = webMenuParams.State
		newCategory.Sort = webMenuParams.Sort
		newCategory.ParentId = webMenuParams.ParentId
		err = webCategoryServer.EditCategory(newCategory)
	}
	if err != nil {
		common.Failed(ctx,"修改菜单失败")
		return
	}
	common.Success(ctx,"ok")
	return
}
// 添加前台菜单
func PostAddWebMenu(ctx *gin.Context) {
	var webMenuParams backend_param.WebMenuParams
	err := ctx.ShouldBindBodyWith(&webMenuParams,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取菜单参数失败")
		return
	}
	if webMenuParams.ParentId == 0 {
		var newMenu web_model.MenuModel
		newMenu.Title = webMenuParams.Title
		newMenu.Sort = webMenuParams.Sort
		newMenu.State = webMenuParams.State
		err = webMenuServer.AddMenu(newMenu)
	} else {
		var newCategory web_model.CategoryModel
		newCategory.Title = webMenuParams.Title
		newCategory.State = webMenuParams.State
		newCategory.Sort = webMenuParams.Sort
		newCategory.ParentId = webMenuParams.ParentId
		err = webCategoryServer.AddCategory(newCategory)
	}
	if err != nil {
		common.Failed(ctx,"添加菜单失败")
		return
	}
	common.Success(ctx,"ok")
	return
}
//  删除菜单
func PostRemoveMenu(ctx *gin.Context) {
	var webMenuParams backend_param.WebMenuParams
	err := ctx.ShouldBindBodyWith(&webMenuParams,binding.JSON)
	if err != nil {
		common.Failed(ctx,"获取菜单参数失败")
		return
	}
	id := webMenuParams.Id
	if webMenuParams.ParentId == 0 {
		err = webMenuServer.RemoveMenu(id)
	} else {
		err = webCategoryServer.RemoveCategory(id)
	}
	if err != nil {
		common.Failed(ctx,"删除菜单失败")
		return
	}
	common.Success(ctx,"ok")
	return
}