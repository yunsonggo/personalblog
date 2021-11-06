package backend_param

import (
	"2021/yunsongcailu/yunsong_server/web/web_model"
	"mime/multipart"
)

type BackendArticleParam struct {
	Count int `json:"count"`
	Start int `json:"start"`
}

type BackendArticleRemoveParam struct {
	IdStrArr string `json:"id_str_arr"`
}

type ArticleBannerUploadParam struct {
	Token string `form:"token"`
	OldBanner string `form:"old_banner_url"`
	ArticleId int64 `form:"article_id"`
	File *multipart.FileHeader `form:"file"`
}

type ArticleCoverUploadParam struct {
	Token string `form:"token"`
	OldCover string `form:"old_cover_url"`
	ArticleId int64 `form:"article_id"`
	File *multipart.FileHeader `form:"file"`
}

type ArticleVideoUploadParam struct {
	Token string `form:"token"`
	OldVideo string `form:"old_video_url"`
	ArticleId int64 `form:"article_id"`
	File *multipart.FileHeader `form:"file"`
}

type ArticleAddParam struct {
	ArticleInfo web_model.ArticleModel `json:"article_info"`
}

type ArticleEditParam struct {
	ArticleId int64 `json:"article_id"`
	ArticleInfo web_model.ArticleModel `json:"article_info"`
}

//type ArticleInfoAddParam struct {
//	Id int64 `json:"id"`
//	ArticleId int64 `json:"article_id"`
//	Title string `json:"title"`
//	Desc string `json:"desc"`
//	Keyword string `json:"keyword"`
//	IsOriginal int `json:"is_original"`
//	LinkUrl string `json:"link_url"`
//	State int `json:"state"`
//	MenuId int64 `json:"menu_id"`
//	CategoryId int64 `json:"category_id"`
//	IsIndex int `json:"is_index"`
//	IsTop int `json:"is_top"`
//	IsBottom int `json:"is_bottom"`
//	Author string `json:"author"`
//	AuthorId int64 `json:"author_id"`
//	Content string `json:"content"`
//	Cover string `json:"cover"`
//	Banner string `json:"banner"`
//	Video string `json:"video"`
//}