package backend_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type BackendLinkDao interface {
	// 获取所有链接
	FindLinkAll() (linkList []web_model.LinksModel, err error)
	// 根据ID 修改链接
	UpdateLinkById(link web_model.LinksModel) (err error)
	// 根据ID 修改链接图片
	UpdateLinkIcon(id int64,path string) (err error)
	// 根据ID 删除链接
	DeleteLinkById(id int64) (err error)
	// 添加链接
	InsertLink(link web_model.LinksModel) (err error)
}

type backendLinkDao struct {}

func NewBackendLinkDao() BackendLinkDao {
	return &backendLinkDao{}
}

// 获取所有链接
func (bld *backendLinkDao) FindLinkAll() (linkList []web_model.LinksModel, err error) {
	err = dial.DB.Find(&linkList)
	return
}
// 根据ID 修改链接
func (bld *backendLinkDao) UpdateLinkById(link web_model.LinksModel) (err error) {
	_,err = dial.DB.Where("id = ?",link.Id).Cols("link_icon","link_url","sort","link_title").Update(&link)
	return
}
// 根据ID 修改链接图片
func (bld *backendLinkDao) UpdateLinkIcon(id int64,path string) (err error) {
	newLink := new(web_model.LinksModel)
	newLink.LinkIcon = path
	_,err = dial.DB.Where("id = ?",id).Cols("link_icon").Update(newLink)
	return
}
// 根据ID 删除链接
func (bld *backendLinkDao) DeleteLinkById(id int64) (err error) {
	newLink := new(web_model.LinksModel)
	_,err = dial.DB.Where("id = ?",id).Delete(newLink)
	return
}
// 添加链接
func (bld *backendLinkDao) InsertLink(link web_model.LinksModel) (err error) {
	_,err = dial.DB.InsertOne(&link)
	return
}