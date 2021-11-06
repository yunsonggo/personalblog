package backend_service

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_dao"
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
)

type BackendCategoryServer interface {
	// 获取所有类别
	GetCategories() (categories []backend_model.BackendCategoryModel,err error)
	// 添加类别
	AddCategory(category backend_model.BackendCategoryModel) (err error)
	// 修改类别
	EditCategory(category backend_model.BackendCategoryModel) (err error)
	// 删除类别
	RemoveCategory(categoryId int64) (err error)
	// 根据类别ID 查询分类
	GetCategoriesByMenuId(menuId int64) (categories []backend_model.BackendCategoryModel,err error)
}

type backendCategoryServer struct {}

func NewBackendCategoryServer() BackendCategoryServer {
	return &backendCategoryServer{}
}

var bcd = backend_dao.NewBackendCategoryDao()

// 获取所有类别
func(bcs *backendCategoryServer) GetCategories() (categories []backend_model.BackendCategoryModel,err error) {
	return bcd.FindCategories()
}
// 添加类别
func(bcs *backendCategoryServer) AddCategory(category backend_model.BackendCategoryModel) (err error) {
	return bcd.InsertCategory(category)
}
// 修改类别
func(bcs *backendCategoryServer) EditCategory(category backend_model.BackendCategoryModel) (err error) {
	return bcd.UpdateCategory(category)
}
// 删除类别
func(bcs *backendCategoryServer) RemoveCategory(categoryId int64) (err error) {
	return bcd.DeleteCategory(categoryId)
}
// 根据类别ID 查询分类
func(bcs *backendCategoryServer) GetCategoriesByMenuId(menuId int64) (categories []backend_model.BackendCategoryModel,err error) {
	return bcd.FindCategoriesByMenuId(menuId)
}