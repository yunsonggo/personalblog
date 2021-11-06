package backend_service

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type ConsumerServer interface {
	// 获取用户列表
	FindConsumers() (consumerList []web_model.Consumers,err error)
	// 修改用户状态  1删除 0 激活
	EditConsumerState(id int64,state int) (err error)
}

type consumerServer struct {}

func NewConsumerServer() ConsumerServer {
	return &consumerServer{}
}
var backendConsumerDao = backend_dao.NewConsumerDao()
// 获取用户列表
func (bcs *consumerServer) FindConsumers() (consumerList []web_model.Consumers,err error) {
	return backendConsumerDao.QueryConsumers()
}
// 修改用户状态  1删除 0 激活
func (bcs *consumerServer) EditConsumerState(id int64,state int) (err error) {
	return backendConsumerDao.UpdateConsumerState(id,state)
}