package web_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type WebsiteDao interface {
	QueryWebsiteInfo() (websiteInfo web_model.WebsiteModel,err error)
}

type websiteDao struct {}

func NewWebsiteDao() WebsiteDao {
	return &websiteDao{}
}

func (wd *websiteDao) QueryWebsiteInfo() (websiteInfo web_model.WebsiteModel,err error) {
	_,err = dial.DB.Where("id = 1").Get(&websiteInfo)
	return
}