package web_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type CategoryDao interface {
	// 获取所有类别
	QueryCategoryAll() (categoryAll []web_model.CategoryModel,err error)
	// 获取所有激活类别
	QueryCategory() (categories []web_model.CategoryModel,err error)
	// 修改类别
	UpdateCategory(category web_model.CategoryModel) (err error)
	// 添加类别
	InsertCategory(category web_model.CategoryModel) (err error)
	// 删除类别
	DeleteCategory(categoryId int64) (err error)
}

type categoryDao struct {}

func NewCategoryDao() CategoryDao {
	return &categoryDao{}
}

// 获取所有类别
func (cgd *categoryDao) QueryCategoryAll() (categoryAll []web_model.CategoryModel,err error) {
	err = dial.DB.Find(&categoryAll)
	return
}
// 获取所有激活类别
func (cgd *categoryDao) QueryCategory() (categories []web_model.CategoryModel,err error) {
	err = dial.DB.Where("state = 0").Find(&categories)
	return
}
// 修改类别
func (cgd *categoryDao) UpdateCategory(category web_model.CategoryModel) (err error) {
	_,err = dial.DB.Where("id = ?",category.Id).Cols("title","sort","state","parent_id").Update(&category)
	return
}
// 添加类别
func (cgd *categoryDao) InsertCategory(category web_model.CategoryModel) (err error) {
	_,err = dial.DB.InsertOne(&category)
	return
}
// 删除类别
func (cgd *categoryDao) DeleteCategory(categoryId int64) (err error) {
	newCategory := new(web_model.CategoryModel)
	_,err = dial.DB.Where("id = ?",categoryId).Delete(newCategory)
	return
}