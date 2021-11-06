package web_service

import (
	"2021/yunsongcailu/yunsong_server/web/web_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type CategoryServer interface {
	// 获取所有类别
	FindCategoryAll() (categoryAll []web_model.CategoryModel,err error)
	// 获取所有激活类别
	FindCategory() (categories []web_model.CategoryModel,err error)
	// 修改类别
	EditCategory(category web_model.CategoryModel) (err error)
	// 添加类别
	AddCategory(category web_model.CategoryModel) (err error)
	// 删除类别
	RemoveCategory(categoryId int64) (err error)
}

type categoryServer struct {}

func NewCategoryServer() CategoryServer {
	return &categoryServer{}
}
var cgd = web_dao.NewCategoryDao()

// 获取所有类别
func (cgs *categoryServer) FindCategoryAll() (categoryAll []web_model.CategoryModel,err error) {
	return cgd.QueryCategoryAll()
}
// 获取所有激活类别
func (cgs *categoryServer) FindCategory() (categories []web_model.CategoryModel,err error) {
	return cgd.QueryCategory()
}
// 修改类别
func (cgs *categoryServer) EditCategory(category web_model.CategoryModel) (err error) {
	return cgd.UpdateCategory(category)
}
// 添加类别
func (cgs *categoryServer) AddCategory(category web_model.CategoryModel) (err error) {
	return cgd.InsertCategory(category)
}
// 删除类别
func (cgs *categoryServer) RemoveCategory(categoryId int64) (err error) {
	return cgd.DeleteCategory(categoryId)
}