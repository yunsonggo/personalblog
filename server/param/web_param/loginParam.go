package web_param

type EmailLoginParam struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Captcha string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}

type PhoneLoginParam struct {
	Phone string `json:"phone" binding:"required"`
	Code string `json:"code" binding:"required"`
	BizId string `json:"bizid" binding:"required"`
}