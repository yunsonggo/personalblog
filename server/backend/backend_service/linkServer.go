package backend_service

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type BackendLinkServer interface {
	// 获取所有链接
	GetLinkAll() (linkList []web_model.LinksModel, err error)
	// 根据ID 修改链接
	EditLinkById(link web_model.LinksModel) (err error)
	// 根据ID 修改链接图片
	EditLinkIcon(id int64,path string) (err error)
	// 根据ID 删除链接
	RemoveLinkById(id int64) (err error)
	// 添加链接
	AddLink(link web_model.LinksModel) (err error)
}

type backendLinkServer struct {}

func NewBackendLinkServer() BackendLinkServer {
	return &backendLinkServer{}
}

var bld = backend_dao.NewBackendLinkDao()

// 获取所有链接
func (bls *backendLinkServer) GetLinkAll() (linkList []web_model.LinksModel, err error) {
	return bld.FindLinkAll()
}
// 根据ID 修改链接
func (bls *backendLinkServer) EditLinkById(link web_model.LinksModel) (err error) {
	return bld.UpdateLinkById(link)
}
// 根据ID 修改链接图片
func (bls *backendLinkServer) EditLinkIcon(id int64,path string) (err error) {
	return bld.UpdateLinkIcon(id,path)
}
// 根据ID 删除链接
func (bls *backendLinkServer) RemoveLinkById(id int64) (err error) {
	return bld.DeleteLinkById(id)
}
// 添加链接
func (bls *backendLinkServer) AddLink(link web_model.LinksModel) (err error) {
	return bld.InsertLink(link)
}