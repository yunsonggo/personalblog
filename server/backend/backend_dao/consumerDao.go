package backend_dao

import (
	"2021/yunsongcailu/yunsong_server/dial"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type ConsumerDao interface {
	// 获取用户列表
	QueryConsumers() (consumerList []web_model.Consumers,err error)
	// 修改用户状态  1删除 0 激活
	UpdateConsumerState(id int64,state int) (err error)
}

type consumerDao struct {}

func NewConsumerDao() ConsumerDao {
	return &consumerDao{}
}

// 获取用户列表
func (bcd *consumerDao) QueryConsumers() (consumerList []web_model.Consumers,err error) {
	err = dial.DB.Find(&consumerList)
	return
}
// 修改用户状态  1删除 0 激活
func (bcd *consumerDao) UpdateConsumerState(id int64,state int) (err error) {
	var newConsumer web_model.Consumers
	newConsumer.IsDel = state
	_,err = dial.DB.Where("id = ?",id).Cols("is_del").Update(&newConsumer)
	return
}