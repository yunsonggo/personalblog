package web_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type ConsumerDao interface {
	// 根据ID 查询用户
	QueryConsumerById(id int64) (consumer web_model.Consumers,err error)
	// 根据邮箱查询用户
	QueryConsumerByEmail(email string) (consumer web_model.Consumers,err error)
	// 根据手机查询用户
	QueryConsumerByPhone(phone string) (consumer web_model.Consumers,err error)
	// 新用户入库
	InsertConsumer(consumer web_model.Consumers) (id int64,err error)
	// 根据ID 更新头像
	UpdateConsumerIcon(id int64,filePath string) (err error)
	// 更新用户信息
	UpdataConsumerInfoById(consumer web_model.Consumers) (err error)
}

type consumerDao struct {}

func NewConsumerDao() ConsumerDao {
	return &consumerDao{}
}

// 根据ID 查询用户
func (cd *consumerDao) QueryConsumerById(id int64) (consumer web_model.Consumers,err error) {
	_,err = dial.DB.Where("id = ?",id).Get(&consumer)
	return
}
// 根据邮箱查询用户
func (cd *consumerDao) QueryConsumerByEmail(email string) (consumer web_model.Consumers,err error) {
	_,err = dial.DB.Where("email = ?",email).Get(&consumer)
	return
}
// 根据手机查询用户
func (cd *consumerDao) QueryConsumerByPhone(phone string) (consumer web_model.Consumers,err error) {
	_,err = dial.DB.Where("phone = ?",phone).Get(&consumer)
	return
}
// 新用户入库
func (cd *consumerDao) InsertConsumer(consumer web_model.Consumers) (id int64,err error) {
	return dial.DB.InsertOne(&consumer)
}
// 根据ID 更新头像
func (cd *consumerDao) UpdateConsumerIcon(id int64,filePath string) (err error) {
	newConsumer := new(web_model.Consumers)
	newConsumer.Icon = filePath
	_,err = dial.DB.Where("id = ?",id).Update(newConsumer)
	return
}
// 更新用户信息
func (cd *consumerDao) UpdataConsumerInfoById(consumer web_model.Consumers) (err error) {
	_,err = dial.DB.Where("id = ?",consumer.Id).Update(consumer)
	return
}