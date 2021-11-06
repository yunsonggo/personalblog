package web_param

type CommentParam struct {
	Id int64 `json:"id"`
	ArticleId int64 `json:"article_id"`
	AuthorId int64  `json:"author_id"`
	AuthorNickname string `json:"author_nickname"`
	AuthorIcon string `json:"author_icon"`
	Content string `json:"content"`
	Count int `json:"count"`
	Start int `json:"start"`
}

type CommentChildParam struct {
	Id int64 `json:"id"`
	CommentId int64 `xorm:"int index" json:"comment_id"`
	AuthorId int64  `json:"author_id"`
	AuthorNickname string `json:"author_nickname"`
	AuthorIcon string `json:"author_icon"`
	Content string `json:"content"`
	Count int `json:"count"`
	Start int `json:"start"`
}