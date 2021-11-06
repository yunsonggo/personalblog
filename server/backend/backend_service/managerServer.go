package backend_service

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_dao"
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
)

type ManagerServer interface {
	// 添加用户
	AddManager(manager backend_model.ManagerModel) (err error)
	// 根据账号密码查询用户
	GetManagerByPassword(name string) (manager backend_model.ManagerModel,err error)
	// 根据ID 查询用户
	GetManagerById(id int64) (manager backend_model.ManagerModel,err error)
}

type managerServer struct {}

func NewManagerServer() ManagerServer {
	return &managerServer{}
}
var md = backend_dao.NewManagerDao()

// 添加用户
func (ms *managerServer) AddManager(manager backend_model.ManagerModel) (err error) {
	return md.InsertManager(manager)
}
// 根据账号密码查询用户
func (ms *managerServer) GetManagerByPassword(name string) (manager backend_model.ManagerModel,err error) {
	return md.QueryManagerByPassword(name)
}
// 根据ID 查询用户
func (ms *managerServer) GetManagerById(id int64) (manager backend_model.ManagerModel,err error) {
	return md.QueryManagerById(id)
}