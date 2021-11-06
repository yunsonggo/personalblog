package web_model

type MenuModel struct {
	Id int64
	Title string `xorm:"varchar(12)" json:"title"`
	Sort int `xorm:"int" json:"sort"`
	State int `xorm:"int" json:"state"` // 0 激活  1 隐藏
}
