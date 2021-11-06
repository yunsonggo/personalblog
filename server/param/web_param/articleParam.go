package web_param

type ArticleParam struct {
	Id         int64  `json:"id"`
	MenuId     int64  `json:"menu_id"`
	CategoryId int64  `json:"category_id"`
	ArticleId  int64  `json:"article_id"`
	IsIndex    int    `json:"is_index"`
	IsTop      int    `json:"is_top"`
	IsBottom   int    `json:"is_bottom"`
	IsOriginal int    `json:"is_original"`
	AuthorId   int64  `json:"author_id"`
	Count      int    `json:"count"`
	Start      int    `json:"start"`
	JobTag     string `json:"job_tag"`
	Star       int    `json:"star"`  // 点赞
	Tread      int    `json:"tread"` // 踩
	SearchKeyword string `json:"keyword"`
}
