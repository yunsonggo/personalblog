package web_param

import "mime/multipart"

type PersonalEditParam struct {
	Consumer PersonalParams `json:"consumer"`
	IsEmail string `json:"is_email"`
	IsPhone string `json:"is_phone"`
}

type PersonalParams struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Name string `json:"name"`
	Nickname string `json:"nickname"`
	Gender string `json:"gender"`
	Icon string `json:"icon"`
	Desc string `json:"desc"`
}

type PersonalAvatarParams struct {
	Token string `form:"token"`
	OldAvatar string `form:"old_avatar"`
	File *multipart.FileHeader `form:"file"`
}