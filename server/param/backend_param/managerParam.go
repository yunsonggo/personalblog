package backend_param

type ManagerParam struct {
	Id int64 `json:"id"`
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
	Captcha string `json:"captcha" binding:"required"`
}
