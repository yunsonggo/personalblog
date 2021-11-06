package backend_service

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_dao"
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
	"2021/yunsongcailu/yunsong_server/web/web_service"
)

type BackendMenuServer interface {
	// 获取所有菜单
	GetMenus() (menus []backend_model.BackendMenuListModel,err error)
	// 添加菜单
	AddMenu(menu backend_model.BackendMenuModel) (err error)
	// 修改菜单
	EditMenu(menu backend_model.BackendMenuModel) (err error)
	// 删除菜单
	RemoveMenu(menuId int64) (err error)
	// 获取所有前台菜单
	FindWebMenuAll() (webMenuLists []backend_model.WebMenuListModel,err error)
}

type backendMenuServer struct {}

func NewBackendMenuServer() BackendMenuServer {
	return &backendMenuServer{}
}

var bmd = backend_dao.NewBackendMenuDao()
var wms = web_service.NewMenuServer()
var wcs = web_service.NewCategoryServer()
// 获取所有菜单
func (bms *backendMenuServer) GetMenus() (menus []backend_model.BackendMenuListModel,err error) {
	menuArr ,err := bmd.FindMenus()
	if err != nil {
		return
	}
	for _,menu := range menuArr {
		var newMenuList backend_model.BackendMenuListModel
		categories,err := bcd.FindCategoriesByMenuId(menu.Id)
		if err != nil {
			continue
		}
		newMenuList.Menu = menu
		newMenuList.MenuCategory = categories
		menus = append(menus, newMenuList)
	}
	return
}
// 添加菜单
func (bms *backendMenuServer) AddMenu(menu backend_model.BackendMenuModel) (err error) {
	return bmd.InsertMenu(menu)
}
// 修改菜单
func (bms *backendMenuServer) EditMenu(menu backend_model.BackendMenuModel) (err error) {
	return bmd.UpdateMenu(menu)
}
// 删除菜单
func (bms *backendMenuServer) RemoveMenu(menuId int64) (err error) {
	return bmd.DeleteMenu(menuId)
}
// 获取所有前台菜单
func (bms *backendMenuServer) FindWebMenuAll() (webMenuLists []backend_model.WebMenuListModel,err error) {
	webMenus ,err := wms.FindMenuAll()
	if err != nil {
		return
	}
	webCategories,err := wcs.FindCategoryAll()
	if err != nil {
		return
	}
	for _,webMenu := range webMenus {
		var webMenuList backend_model.WebMenuListModel
		webMenuList.WebMenu = webMenu
		for _,webCategory := range webCategories {
			if webCategory.ParentId == webMenu.Id {
				webMenuList.WebMenuCategories = append(webMenuList.WebMenuCategories, webCategory)
			}
		}
		webMenuLists = append(webMenuLists, webMenuList)
	}
	return
}