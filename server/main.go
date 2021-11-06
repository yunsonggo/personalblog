package main

import (
	"2021/yunsongcailu/yunsong_server/config"
	"2021/yunsongcailu/yunsong_server/router"
	"github.com/yakaa/log4g"
)

func main() {
	// 初始化配置文件
	config.ConfigInit()
	// 启用数据库
	config.InitRrsEngine()
	// 启用路由
	app := router.NewRouter()
	// 启用服务
	if config.Conf.Website.WebsiteMode == "release" {
		log4g.Error(app.Run(config.Conf.Listen))
	} else {
		_ = app.Run(config.Conf.Listen)
	}
}
