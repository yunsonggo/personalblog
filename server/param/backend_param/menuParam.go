package backend_param

// 菜单参数
type WebMenuParams struct {
	Id int64 `json:"Id"`
	Title string `json:"title"`
	Sort int `json:"sort"`
	State int `json:"state"`
	ParentId int64 `json:"parent_id"`
}
