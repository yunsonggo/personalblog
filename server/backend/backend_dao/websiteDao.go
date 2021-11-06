package backend_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type BackendWebsiteDao interface {
	// 获取网站数据
	FindWebsite() (website web_model.WebsiteModel,err error)
	// 更新网站头像图片
	UpdateAvatar(id int64,avatarPath string) (err error)
	// 更新网站logo
	UpdateLogo(id int64,logoPath string) (err error)
	// 更新网站banner
	UpdateBanner(id int64,bannerPath string) (err error)
	// 更新网站微博
	UpdateWebo(id int64,weboPath string) (err error)
	// 更新网站微信
	UpdateWechar(id int64,wecharPath string) (err error)
	// 更新网站基本信息
	UpdateWebsiteInfo(website web_model.WebsiteModel) (err error)
	// 更新字段
	UpdateWebsiteText(field,value string) (err error)
}

type backendWebsiteDao struct {}

func NewBackendWebsiteDao() BackendWebsiteDao {
	return &backendWebsiteDao{}
}
// 获取网站数据
func (bwd *backendWebsiteDao) FindWebsite() (website web_model.WebsiteModel,err error) {
	_,err = dial.DB.Where("id = 1").Get(&website)
	return
}
// 更新网站头像图片
func (bwd *backendWebsiteDao) UpdateAvatar(id int64,avatarPath string) (err error) {
	newWebsite := new(web_model.WebsiteModel)
	newWebsite.Id = id
	newWebsite.WebsiteAvatar = avatarPath
	_,err = dial.DB.Where("id = ?",id).Cols("website_avatar").Update(newWebsite)
	return
}
// 更新网站logo
func (bwd *backendWebsiteDao) UpdateLogo(id int64,logoPath string) (err error) {
	newWebsite := new(web_model.WebsiteModel)
	newWebsite.Id = id
	newWebsite.WebsiteLogo = logoPath
	_,err = dial.DB.Where("id = ?",id).Cols("website_logo").Update(newWebsite)
	return
}
// 更新网站banner
func (bwd *backendWebsiteDao) UpdateBanner(id int64,bannerPath string) (err error) {
	newWebsite := new(web_model.WebsiteModel)
	newWebsite.Id = id
	newWebsite.WebsiteBanner = bannerPath
	_,err = dial.DB.Where("id = ?",id).Cols("website_banner").Update(newWebsite)
	return
}
// 更新网站微博
func (bwd *backendWebsiteDao) UpdateWebo(id int64,weboPath string) (err error) {
	newWebsite := new(web_model.WebsiteModel)
	newWebsite.Id = id
	newWebsite.WebsiteWebo = weboPath
	_,err = dial.DB.Where("id = ?",id).Cols("website_webo").Update(newWebsite)
	return
}
// 更新网站微信
func (bwd *backendWebsiteDao) UpdateWechar(id int64,wecharPath string) (err error) {
	newWebsite := new(web_model.WebsiteModel)
	newWebsite.Id = id
	newWebsite.WebsiteWechar = wecharPath
	_,err = dial.DB.Where("id = ?",id).Cols("website_wechar").Update(newWebsite)
	return
}
// 更新网站基本信息
func (bwd *backendWebsiteDao) UpdateWebsiteInfo(website web_model.WebsiteModel) (err error) {
	_,err = dial.DB.Where("id = 1").Update(&website)
	return
}
// 更新字段
func (bwd *backendWebsiteDao) UpdateWebsiteText(field,value string) (err error) {
	newWebsite := new(web_model.WebsiteModel)
	switch field {
	case "about":
		newWebsite.WebsiteAbout = value
		_,err = dial.DB.Where("id = 1").Cols("website_about").Update(newWebsite)
	case "agreement":
		newWebsite.WebsiteAgreement = value
		_,err = dial.DB.Where("id = 1").Cols("website_agreement").Update(newWebsite)
	case "privacy":
		newWebsite.WebsitePrivacy = value
		_,err = dial.DB.Where("id = 1").Cols("website_privacy").Update(newWebsite)
	case "apk":
		newWebsite.WebsiteApk = value
		_,err = dial.DB.Where("id = 1").Cols("website_apk").Update(newWebsite)
	}
	return
}