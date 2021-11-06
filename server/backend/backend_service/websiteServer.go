package backend_service

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type BackendWebsiteServer interface {
	// 获取网站数据
	GetWebsite() (website web_model.WebsiteModel,err error)
	// 更新网站头像图片
	EditAvatar(id int64,avatarPath string) (err error)
	// 更新网站logo
	EditLogo(id int64,logoPath string) (err error)
	// 更新网站banner
	EditBanner(id int64,bannerPath string) (err error)
	// 更新网站微博
	EditWebo(id int64,weboPath string) (err error)
	// 更新网站微信
	EditWechar(id int64,wecharPath string) (err error)
	// 更新网站基本信息
	EditWebsiteInfo(website web_model.WebsiteModel) (err error)
	// 更新字段
	EditWebsiteText(field,value string) (err error)
}

type backendWebsiteServer struct {}

func NewBackendWebsiteServer() BackendWebsiteServer {
	return &backendWebsiteServer{}
}

var backendWebsiteDao = backend_dao.NewBackendWebsiteDao()

// 获取网站数据
func (bcs *backendWebsiteServer) GetWebsite() (website web_model.WebsiteModel,err error) {
	return backendWebsiteDao.FindWebsite()
}
// 更新网站头像图片
func (bcs *backendWebsiteServer) EditAvatar(id int64,avatarPath string) (err error) {
	return backendWebsiteDao.UpdateAvatar(id,avatarPath)
}
// 更新网站logo
func (bcs *backendWebsiteServer) EditLogo(id int64,logoPath string) (err error) {
	return backendWebsiteDao.UpdateLogo(id,logoPath)
}
// 更新网站banner
func (bcs *backendWebsiteServer) EditBanner(id int64,bannerPath string) (err error) {
	return backendWebsiteDao.UpdateBanner(id,bannerPath)
}
// 更新网站微博
func (bcs *backendWebsiteServer) EditWebo(id int64,weboPath string) (err error) {
	return backendWebsiteDao.UpdateWebo(id,weboPath)
}
// 更新网站微信
func (bcs *backendWebsiteServer) EditWechar(id int64,wecharPath string) (err error) {
	return backendWebsiteDao.UpdateWechar(id,wecharPath)
}
// 更新网站基本信息
func (bcs *backendWebsiteServer) EditWebsiteInfo(website web_model.WebsiteModel) (err error) {
	return backendWebsiteDao.UpdateWebsiteInfo(website)
}
// 更新字段
func (bcs *backendWebsiteServer) EditWebsiteText(field,value string) (err error) {
	return backendWebsiteDao.UpdateWebsiteText(field,value)
}