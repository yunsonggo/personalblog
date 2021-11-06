package backend_param

import (
	"2021/yunsongcailu/yunsong_server/web/web_model"
	"mime/multipart"
)

type WebsiteAvatarParam struct {
	Token string `form:"token"`
	OldAvatar string `form:"old_avatar_url"`
	Id int64 `form:"id"`
	File *multipart.FileHeader `form:"file"`
}

type WebsiteLogoParam struct {
	Token string `form:"token"`
	OldLogo string `form:"old_logo_url"`
	Id int64 `form:"id"`
	File *multipart.FileHeader `form:"file"`
}

type WebsiteBannerParam struct {
	Token string `form:"token"`
	OldBanner string `form:"old_banner_url"`
	Id int64 `form:"id"`
	File *multipart.FileHeader `form:"file"`
}

type WebsiteWeboParam struct {
	Token string `form:"token"`
	OldWebo string `form:"old_webo_url"`
	Id int64 `form:"id"`
	File *multipart.FileHeader `form:"file"`
}

type WebsiteWecharParam struct {
	Token string `form:"token"`
	OldWechar string `form:"old_wechar_url"`
	Id int64 `form:"id"`
	File *multipart.FileHeader `form:"file"`
}

type WebsiteApkParam struct {
	Token string `form:"token"`
	OldApk string `form:"old_apk_url"`
	Id int64 `form:"id"`
	File *multipart.FileHeader `form:"file"`
}

type WebsiteIosParam struct {
	Token string `form:"token"`
	OldIos string `form:"old_ios_url"`
	Id int64 `form:"id"`
	File *multipart.FileHeader `form:"file"`
}

type WebsiteParam struct {
	Website web_model.WebsiteModel `json:"website_info"`
}
type AboutParam struct {
	About string `json:"about"`
}
type AgreementParam struct {
	Agreement string `json:"agreement"`
}
type PrivacyParam struct {
	Privacy string `json:"privacy"`
}