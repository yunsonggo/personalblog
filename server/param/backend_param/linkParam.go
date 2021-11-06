package backend_param

import "mime/multipart"

type LinkTextParam struct {
	Id int64 `json:"id"`
	Icon string `json:"icon"`
	LinkUrl string `json:"link_url"`
	Sort int `json:"sort"`
	Title string `json:"title"`
}

type LinkIconUploadParam struct {
	Token string `form:"token"`
	OldIcon string `form:"old_icon_url"`
	LinkId int64 `form:"link_id"`
	File *multipart.FileHeader `form:"file"`
}