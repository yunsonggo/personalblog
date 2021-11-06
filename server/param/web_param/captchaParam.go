package web_param

// 消费者 获取邮箱验证码参数
type EmailParam struct {
	Email string `json:"consumer_email" binding:"required"`
	Code string `json:"consumer_email_code"`
}
// 消费者 获取手机验证码参数
type PhoneParam struct {
	Phone string `json:"consumer_phone" binding:"required"`
	Code string `json:"consumer_phone_code"`
	TplId string `json:"tpl_id" binding:"required"`
	Key string `json:"key" binding:"required"`
}