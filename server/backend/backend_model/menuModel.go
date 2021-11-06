package backend_model

import "2021/yunsongcailu/yunsong_server/web/web_model"

type BackendMenuModel struct {
	Id int64
	MenuTitle string `xorm:"varchar(30)" json:"menu_title"`
	MenuIcon string `xorm:"varchar(50)" json:"menu_icon"`
	MenuSort int `xorm:"int" json:"menu_sort"`
	MenuPath string `xorm:"varchar(30)" json:"menu_path"`
}

type BackendCategoryModel struct {
	Id int64
	MenuId int64 `xorm:"int index" json:"menu_id"`
	CategoryTitle string `xorm:"varchar(30)" json:"category_title"`
	CategoryIcon string `xorm:"varchar(50)" json:"category_icon"`
	CategorySort int `xorm:"int" json:"category_sort"`
	CategoryPath string `xorm:"varchar(50)" json:"category_path"`
}

type BackendMenuListModel struct {
	Menu BackendMenuModel
	MenuCategory []BackendCategoryModel
}

type WebMenuListModel struct {
	WebMenu web_model.MenuModel
	WebMenuCategories []web_model.CategoryModel
}