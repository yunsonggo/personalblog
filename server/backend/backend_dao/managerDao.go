package backend_dao

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
	"2021/yunsongcailu/yunsong_server/dial"
)

type ManagerDao interface {
	// 添加用户
	InsertManager(manager backend_model.ManagerModel) (err error)
	// 根据账号密码查询用户
	QueryManagerByPassword(name string) (manager backend_model.ManagerModel,err error)
	// 根据ID 查询用户
	QueryManagerById(id int64) (manager backend_model.ManagerModel,err error)
}

type managerDao struct {}

func NewManagerDao() ManagerDao {
	return &managerDao{}
}

// 添加用户
func (md *managerDao) InsertManager(manager backend_model.ManagerModel) (err error) {
	_,err = dial.DB.InsertOne(&manager)
	return
}
// 根据账号密码查询用户
func (md *managerDao) QueryManagerByPassword(name string) (manager backend_model.ManagerModel,err error) {
	_,err = dial.DB.Where("manager_name = ?",name).Get(&manager)
	return
}
// 根据ID 查询用户
func (md *managerDao) QueryManagerById(id int64) (manager backend_model.ManagerModel,err error) {
	_,err = dial.DB.Where("id = ?",id).Get(&manager)
	return
}