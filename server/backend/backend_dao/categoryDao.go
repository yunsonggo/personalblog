package backend_dao

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
	"2021/yunsongcailu/yunsong_server/dial"
)

type BackendCategoryDao interface {
	// 获取所有类别
	FindCategories() (categories []backend_model.BackendCategoryModel,err error)
	// 添加类别
	InsertCategory(category backend_model.BackendCategoryModel) (err error)
	// 修改类别
	UpdateCategory(category backend_model.BackendCategoryModel) (err error)
	// 删除类别
	DeleteCategory(categoryId int64) (err error)
	// 根据类别ID 查询分类
	FindCategoriesByMenuId(menuId int64) (categories []backend_model.BackendCategoryModel,err error)
}

type backendCategoryDao struct {}

func NewBackendCategoryDao() BackendCategoryDao {
	return &backendCategoryDao{}
}

// 获取所有类别
func (bcd *backendCategoryDao) FindCategories() (categories []backend_model.BackendCategoryModel,err error) {
	err = dial.DB.Find(&categories)
	return
}
// 添加类别
func (bcd *backendCategoryDao) InsertCategory(category backend_model.BackendCategoryModel) (err error) {
	_,err = dial.DB.InsertOne(&category)
	return
}
// 修改类别
func (bcd *backendCategoryDao) UpdateCategory(category backend_model.BackendCategoryModel) (err error) {
	_,err = dial.DB.Where("id = ?",category.Id).Update(&category)
	return
}
// 删除类别
func (bcd *backendCategoryDao) DeleteCategory(categoryId int64) (err error) {
	newCategory := new(backend_model.BackendCategoryModel)
	_,err = dial.DB.Where("id = ?",categoryId).Delete(newCategory)
	return
}
// 根据类别ID 查询分类
func (bcd *backendCategoryDao) FindCategoriesByMenuId(menuId int64) (categories []backend_model.BackendCategoryModel,err error) {
	err = dial.DB.Where("menu_id = ?",menuId).Find(&categories)
	return
}