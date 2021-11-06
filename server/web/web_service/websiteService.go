package web_service

import (
	"2021/yunsongcailu/yunsong_server/web/web_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type WebsiteServer interface {
	GetWebsiteInfo() (websiteInfo web_model.WebsiteModel,err error)
}

type websiteServer struct {}

func NewWebsiteServer() WebsiteServer {
	return &websiteServer{}
}

var wd = web_dao.NewWebsiteDao()

func (ws *websiteServer) GetWebsiteInfo() (websiteInfo web_model.WebsiteModel,err error) {
	return wd.QueryWebsiteInfo()
}