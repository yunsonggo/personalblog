package web_service

import (
	"2021/yunsongcailu/yunsong_server/web/web_dao"
	"2021/yunsongcailu/yunsong_server/web/web_model"
)

type ArticleServer interface {
	// 获取第一个大类的文章 最先展示
	FindIndexArticle(count,start int) (articles []web_model.ArticleModel,err error)
	// 获取某大类首页展示文章
	FindMenuIndexArticle(count,start int) (articles []MenusIndexArticles,err error)
	// 获取指定大类和小类文章
	FindMenuCategoryArticle(menuId,categoryId int64,count,start int) (articles []web_model.ArticleModel,err error)
	// 获取指定大类无小类文章
	FindMenuArticle(menuId int64,count,start int) (articles []web_model.ArticleModel,err error)
	// 根据	ID 获取文章
	GetArticleById(articleId int64) (article web_model.ArticleModel,err error)
	// 根据ID 更新 文章某一字段
	EditArticleOnlyValueById(articleId int64,fieldName string,newArticle *web_model.ArticleModel) (err error)
	// 根据关键字 搜索文章
	FindArticleByKeyword(key,keyword string,count,start int) (articleList []web_model.ArticleModel,err error)
}

type articleServer struct {}

type MenusIndexArticles struct {
	MenuId int64
	MenuTitle string
	Articles []web_model.ArticleModel
}

func NewArticleServer() ArticleServer {
	return &articleServer{}
}

var ad = web_dao.NewArticleDao()

// 获取第一个大类的文章 最先展示
func (as *articleServer) FindIndexArticle(count,start int) (articles []web_model.ArticleModel,err error) {
	// 获取第一大类ID
	menus,err := md.QueryMenu()
	if err != nil {
		return
	}
	menuId := menus[0].Id
	return ad.QueryIndexArticle(menuId,count,start)
}
// 获取各大类首页展示文章
func (as *articleServer) FindMenuIndexArticle(count,start int) (articles []MenusIndexArticles,err error) {
	menus,err := md.QueryMenu()
	if err != nil {
		return
	}
	for key,menu := range menus {
		if key != 0 {
			var menuIndexArticle MenusIndexArticles
			menuIndexArticle.MenuId = menu.Id
			menuIndexArticle.MenuTitle = menu.Title
			res,err := ad.QueryMenuIndexArticle(menu.Id,count,start)
			if err != nil {
				continue
			}
			menuIndexArticle.Articles = res
			articles = append(articles, menuIndexArticle)
		}
	}
	return
}
// 获取指定大类和小类文章
func (as *articleServer) FindMenuCategoryArticle(menuId,categoryId int64,count,start int) (articles []web_model.ArticleModel,err error) {
	start = (start - 1 ) * count + 1
	return ad.QueryMenuCategoryArticle(menuId,categoryId,count,start)
}
// 获取指定大类无小类文章
func (as *articleServer) FindMenuArticle(menuId int64,count,start int) (articles []web_model.ArticleModel,err error) {
	start = (start - 1 ) * count + 1
	return ad.QueryMenuArticle(menuId,count,start)
}
// 根据	ID 获取文章
func (as *articleServer) GetArticleById(articleId int64) (article web_model.ArticleModel,err error) {
	article,err = ad.QueryArticleById(articleId)
	if err != nil {
		return
	}
	article.ReadNum += 1
	newArticle := new(web_model.ArticleModel)
	newArticle.ReadNum = article.ReadNum
	err = as.EditArticleOnlyValueById(article.Id,"read_num",newArticle)
	return
}
// 根据ID 更新 文章某一字段
func (as *articleServer) EditArticleOnlyValueById(articleId int64,fieldName string,newArticle *web_model.ArticleModel) (err error) {
	return ad.UpdateArticleOnlyValueById(articleId,fieldName,newArticle)
}
// 根据关键字 搜索文章
func (as *articleServer) FindArticleByKeyword(key,keyword string,count,start int) (articleList []web_model.ArticleModel,err error) {
	start = (start - 1 ) * count + 1
	return ad.QueryArticleByKeyword(key,keyword,count,start)
}