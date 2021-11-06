package web_service

import (
	"2021/yunsongcailu/yunsong_server/web/web_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type MenuServer interface {
	// 获取所有菜单
	FindMenuAll() (menus []web_model.MenuModel,err error)
	// 获取所有激活菜单
	FindMenu() (menus []web_model.MenuModel,err error)
	// 修改菜单
	EditMenu(menu web_model.MenuModel) (err error)
	// 添加菜单
	AddMenu(menu web_model.MenuModel) (err error)
	// 删除菜单
	RemoveMenu(menuId int64) (err error)
}

type menuServer struct {}

var md = web_dao.NewMenuDao()

func NewMenuServer() MenuServer {
	return &menuServer{}
}

// 获取所有菜单及类别
func (ms *menuServer) FindMenuAll() (menus []web_model.MenuModel,err error) {
	return md.QueryMenuAll()
}
// 获取所有激活菜单
func (ms *menuServer) FindMenu() (menus []web_model.MenuModel,err error) {
	return md.QueryMenu()
}
// 修改菜单
func (ms *menuServer) EditMenu(menu web_model.MenuModel) (err error) {
	return md.UpdateMenu(menu)
}
// 添加菜单
func (ms *menuServer) AddMenu(menu web_model.MenuModel) (err error) {
	return md.InsertMenu(menu)
}
// 删除菜单
func (ms *menuServer) RemoveMenu(menuId int64) (err error) {
	return md.DeleteMenu(menuId)
}