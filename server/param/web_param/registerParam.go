package web_param

type RegisterEmailParam struct {
	Name string `json:"name" building:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	CheckPass string `json:"checkpass" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type RegisterPhoneParam struct {
	Name string `json:"name" building:"required"`
	Phone string `json:"phone" binding:"required"`
	BizId string `json:"bizid" binding:"required"`
	Code string `json:"code" binding:"required"`
}