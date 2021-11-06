package web_model

type CategoryModel struct {
	Id int64
	ParentId int64 `xorm:"int" json:"parent_id"`
	Title string `xorm:"varchar(20)" json:"title"`
	Sort int `xorm:"int" json:"sort"`
	State int `xorm:"int" json:"state"` // 0 激活 1 隐藏
}
