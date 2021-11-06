package web_service

import (
	"2021/yunsongcailu/yunsong_server/tools"
	"2021/yunsongcailu/yunsong_server/web/web_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
	"fmt"
)

type ConsumerServer interface {
	// 根据ID 查询用户
	GetConsumerById(id int64) (consumer web_model.Consumers,err error)
	// 根据邮箱查询用户
	GetConsumerByEmail(email string) (consumer web_model.Consumers,err error)
	// 根据手机查询用户
	GetConsumerByPhone(phone string) (consumer web_model.Consumers,err error)
	// 新用户入库
	AddConsumer(name,email,password,phone string) (*web_model.Consumers,error)
	// 更新用户头像
	EditConsumerIconById(id int64,filePath string) (err error)
	// 更新用户信息
	EditConsumerInfoById(consumer web_model.Consumers) (err error)
}

type consumerService struct {}

func NewConsumerService () ConsumerServer {
	return &consumerService{}
}

var cd = web_dao.NewConsumerDao()

// 根据ID 查询用户
func (cs *consumerService) GetConsumerById(id int64) (consumer web_model.Consumers,err error) {
	return cd.QueryConsumerById(id)
}
// 根据邮箱查询用户
func (cs *consumerService) GetConsumerByEmail(email string) (consumer web_model.Consumers,err error) {
	return cd.QueryConsumerByEmail(email)
}
// 根据手机查询用户
func (cs *consumerService) GetConsumerByPhone(phone string) (consumer web_model.Consumers,err error) {
	return cd.QueryConsumerByPhone(phone)
}
// 新用户入库
func (cs *consumerService) AddConsumer(name,email,password,phone string) (*web_model.Consumers,error) {
	var consumer web_model.Consumers
	consumer.Name = name
	consumer.Email = email
	consumer.Phone = phone
	consumer.Password = tools.EncodeSha256(password)
	consumer.IsDel = 0
	id,err := cd.InsertConsumer(consumer)
	fmt.Printf("返回的id = %d\n",id)
	if err != nil || id == 0 {
		return nil,err
	}
	consumer.Id = id
	return &consumer,nil
}
// 更新用户头像
func (cs *consumerService) EditConsumerIconById(id int64,filePath string) (err error) {
	return cd.UpdateConsumerIcon(id,filePath)
}
// 更新用户信息
func (cs *consumerService) EditConsumerInfoById(consumer web_model.Consumers) (err error) {
	return cd.UpdataConsumerInfoById(consumer)
}