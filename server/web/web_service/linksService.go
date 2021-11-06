package web_service

import (
	"2021/yunsongcailu/yunsong_server/web/web_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type LinksServer interface {
	// 查询所有链接
	FindLinks() (links []web_model.LinksModel,err error)
}

type linksServer struct {}

func NewLinksServer() LinksServer {
	return &linksServer{}
}
var ld = web_dao.NewLinksDao()
// 查询所有链接
func (ls *linksServer) FindLinks() (links []web_model.LinksModel,err error) {
	return ld.QueryLinks()
}