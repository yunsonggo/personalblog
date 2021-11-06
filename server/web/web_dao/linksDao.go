package web_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type LinksDao interface {
	// 查询所有链接
	QueryLinks() (links []web_model.LinksModel,err error)
}

type linksDao struct {}

func NewLinksDao() LinksDao {
	return &linksDao{}
}

// 查询所有链接
func (ld *linksDao) QueryLinks() (links []web_model.LinksModel,err error) {
	err = dial.DB.Find(&links)
	return
}