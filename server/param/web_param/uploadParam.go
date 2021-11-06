package web_param

type UploadAvatarParams struct {
	Token string `json:"token" form:"token"`
	OldAvatar string `json:"old_avatar" form:"old_avatar"`
}
