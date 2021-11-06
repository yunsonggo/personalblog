package backend_model

import (
	"2021/yunsongcailu/yunsong_server/tools"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type BackendComment struct {
	Id int64
	AuthorId int64
	AuthorNickname string
	AuthorIcon string
	ArticleId int64
	Content string
	CreateTime tools.JsonTime
	HasChild bool
	CommentArticle web_model.ArticleModel
	CommentChildren [] BackendCommentChild
}

type BackendCommentChild struct {
	Id int64
	AuthorId int64
	AuthorNickname string
	AuthorIcon string
	CommentId int64
	Content string
	CreateTime tools.JsonTime
}