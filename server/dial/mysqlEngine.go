package dial

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_model"
	"2021/yunsongcailu/yunsong_server/web/web_model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

func MysqlEngine(conn string,driver string,show bool) {
	db,err := xorm.NewEngine(driver,conn)
	if err != nil {
		panic(fmt.Sprintf("连接数据库异常:%v\n",err))
	}
	db.ShowSQL(show)
	err = db.Sync2(
		new(web_model.Consumers),
		new(web_model.MenuModel),
		new(web_model.CategoryModel),
		new(web_model.ArticleModel),
		new(web_model.CommentModel),
		new(web_model.CommentChildModel),
		new(web_model.WebsiteModel),
		new(web_model.LinksModel),
		new(backend_model.ManagerModel),
		new(backend_model.BackendMenuModel),
		new(backend_model.BackendCategoryModel),
		)
	if err != nil {
		panic(fmt.Sprintf("创建数据表出现异常:%v\n",err))
	}
	DB = db
	return
}

