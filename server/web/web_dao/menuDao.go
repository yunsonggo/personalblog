package web_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type MenuDao interface {
	// 获取所有菜单
	QueryMenuAll() (menus []web_model.MenuModel,err error)
	// 获取所有激活菜单
	QueryMenu() (menus []web_model.MenuModel,err error)
	// 修改菜单
	UpdateMenu(menu web_model.MenuModel) (err error)
	// 添加菜单
	InsertMenu(menu web_model.MenuModel) (err error)
	// 删除菜单
	DeleteMenu(menuId int64) (err error)
}

type menuDao struct {}

func NewMenuDao() MenuDao {
	return &menuDao{}
}

// 获取所有菜单
func (md *menuDao) QueryMenuAll() (menus []web_model.MenuModel,err error) {
	err = dial.DB.Find(&menus)
	return
}
// 获取所有激活菜单
func (md *menuDao) QueryMenu() (menus []web_model.MenuModel,err error) {
	err = dial.DB.Where("state = 0").Find(&menus)
	return
}
// 修改菜单
func (md *menuDao) UpdateMenu(menu web_model.MenuModel) (err error) {
	_,err = dial.DB.Where("id = ?",menu.Id).Cols("title","sort","state").Update(&menu)
	return
}
// 添加菜单
func (md *menuDao) InsertMenu(menu web_model.MenuModel) (err error) {
	_,err = dial.DB.InsertOne(&menu)
	return
}
// 删除菜单
func (md *menuDao) DeleteMenu(menuId int64) (err error) {
	newMenu := new(web_model.MenuModel)
	_,err = dial.DB.Where("id = ?",menuId).Delete(newMenu)
	return
}