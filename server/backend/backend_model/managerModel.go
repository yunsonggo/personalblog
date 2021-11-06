package backend_model

type ManagerModel struct {
	Id int64
	ManagerName string `xorm:"varchar(55)" json:"manager_name"`
	ManagerPassword string `xorm:"varchar(255)" json:"manager_password"`
	ManagerPower int `xorm:"int" json:"manager_power"`
}
