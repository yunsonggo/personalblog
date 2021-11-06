package web_model

type LinksModel struct {
	Id int64
	LinkTitle string `xorm:"varchar(255)" json:"title"`
	LinkUrl string `xorm:"varchar(255)" json:"link_url"`
	LinkIcon string `xorm:"varchar(255)" json:"icon"`
	Sort int `xorm:"int" json:"sort"`
}