package backend_dao

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
	"2021/yunsongcailu/yunsong_server/dial"
)

type BackendMenuDao interface {
	// 获取所有菜单
	FindMenus() (menus []backend_model.BackendMenuModel,err error)
	// 添加菜单
	InsertMenu(menu backend_model.BackendMenuModel) (err error)
	// 修改菜单
	UpdateMenu(menu backend_model.BackendMenuModel) (err error)
	// 删除菜单
	DeleteMenu(menuId int64) (err error)
}

type backendMenuDao struct {}

func NewBackendMenuDao() BackendMenuDao {
	return &backendMenuDao{}
}

// 获取所有菜单
func (bmd *backendMenuDao) FindMenus() (menus []backend_model.BackendMenuModel,err error) {
	err = dial.DB.Find(&menus)
	return
}
// 添加菜单
func (bmd *backendMenuDao) InsertMenu(menu backend_model.BackendMenuModel) (err error) {
	_,err = dial.DB.InsertOne(&menu)
	return
}
// 修改菜单
func (bmd *backendMenuDao) UpdateMenu(menu backend_model.BackendMenuModel) (err error) {
	_,err = dial.DB.Where("id = ?",menu.Id).Update(&menu)
	return
}
// 删除菜单
func (bmd *backendMenuDao) DeleteMenu(menuId int64) (err error) {
	menu := new(backend_model.BackendMenuModel)
	_,err = dial.DB.Where("id = ?",menuId).Delete(menu)
	return
}