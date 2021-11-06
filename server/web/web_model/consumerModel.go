package web_model

import "2021/yunsongcailu/yunsong_server/tools"

type Consumers struct {
	Id          int64          `json:"id"`
	Name        string         `xorm:"varchar(20)" json:"name"`
	NickName    string         `xorm:"varchar(20)" json:"nickname"`
	Phone       string         `xorm:"varchar(20)" json:"phone"`
	Email       string         `xorm:"varchar(50)" json:"email"`
	Password    string         `xorm:"varchar(255)" json:"password"`
	Icon        string         `xorm:"varchar(255)" json:"icon"`
	Gender      string         `xorm:"varchar(10)" json:"gender"`
	Desc        string         `xorm:"varchar(255)" json:"desc"`
	IsDel       int            `xorm:"int" json:"is_del"` // 1删除 0 激活
	CreatedTime tools.JsonTime `xorm:"created" json:"created_time"`
}
