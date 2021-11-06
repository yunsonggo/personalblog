package web_model

import "2021/yunsongcailu/yunsong_server/tools"

type CommentModel struct {
	Id int64
	AuthorId int64  `xorm:"int index" json:"author_id"`
	AuthorNickname string `xorm:"varchar(20)" json:"author_nickname"`
	AuthorIcon string `xorm:"varchar(255)" json:"author_icon"`
	ArticleId int64 `xorm:"int index" json:"article_id"`
	Content string `xorm:"text" json:"content"`
	CreateTime tools.JsonTime `xorm:"created" json:"create_time"`  // 时间
}

type CommentChildModel struct {
	Id int64
	AuthorId int64  `xorm:"int index" json:"author_id"`
	AuthorNickname string `xorm:"varchar(20)" json:"author_nickname"`
	AuthorIcon string `xorm:"varchar(255)" json:"author_icon"`
	CommentId int64 `xorm:"int index" json:"comment_id"`
	Content string `xorm:"text" json:"content"`
	CreateTime tools.JsonTime `xorm:"created" json:"create_time"`  // 时间
}