package router

import (
	"2021/yunsongcailu/yunsong_server/backend/backend_controller"
	"2021/yunsongcailu/yunsong_server/common"
	"2021/yunsongcailu/yunsong_server/config"
	"2021/yunsongcailu/yunsong_server/middleware"
	"2021/yunsongcailu/yunsong_server/web/web_controller"
	"github.com/gin-gonic/gin"
	"github.com/yakaa/log4g"
	"net/http"
)

func NewRouter() *gin.Engine {
	wc := config.Conf.Website
	if wc.WebsiteMode == gin.ReleaseMode {
		log4g.Init(log4g.Config{Path: wc.WebsiteLog,Stdout: true})
		gin.DefaultWriter = log4g.InfoLog
		gin.DefaultErrorWriter = log4g.ErrorLog
	}
	// 初始化引擎
	app := gin.Default()
	// 设置静态文件路径
	app.StaticFS("/api/static",http.Dir("./public"))
	// 启用session
	common.InitSession(app)
	// 启用cors
	app.Use(middleware.CorsMiddleware())
	// web接口
	// 自由访问
	webGroup := app.Group("/api")
	{
		// 获取邮箱验证码
		webGroup.POST("/captcha/email",web_controller.PostEmailCaptcha)
		// 获取手机验证码
		webGroup.POST("/captcha/phone",web_controller.PostPhoneCaptcha)
		// 邮箱注册
		webGroup.POST("/register/email",web_controller.PostRegisterByEmail)
		// 手机注册
		webGroup.POST("/register/phone",web_controller.PostRegisterByPhone)
		// 图形验证码
		webGroup.GET("/captcha/img",web_controller.GetImgCaptcha)
		// 邮箱登录
		webGroup.POST("/login/email",web_controller.PostLoginByEmail)
		// 手机登录
		webGroup.POST("/login/phone",web_controller.PostLoginByPhone)
		// 获取所有激活菜单
		webGroup.POST("/menus",web_controller.PostMenu)
		// 获取所有激活类别
		webGroup.POST("/category",web_controller.PostFindCategory)
		// 获取第一大类文章
		webGroup.POST("/article/index",web_controller.PostArticleIndex)
		// 获取非第一大类之外的 类别标记首页展示的文章
		webGroup.POST("/article/other/index",web_controller.PostOtherArticleIndex)
		// 获取指定类别文章
		webGroup.POST("/article/menu/category",web_controller.PostMenuIdCategoryIdArticle)
		// 获取指定大类无小类文章
		webGroup.POST("/article/menu",web_controller.PostMenuArticle)
		// 根据ID 获取文章
		webGroup.POST("/article/info",web_controller.PostArticleById)
		// 根据文章ID 获取评论
		webGroup.POST("/article/comment/list",web_controller.PostCommentByArticleId)
		// 根据文章ID 添加评论
		webGroup.POST("/article/comment/insert",web_controller.PostAddComment)
		// 添加二级评论
		webGroup.POST("/article/comment/insert/child",web_controller.PostAddCommentChild)
		// 更新赞或者踩
		webGroup.POST("/article/good/or",web_controller.PostArticleGoodOrBad)
		// 搜索文章
		webGroup.POST("/article/search/keyword",web_controller.PostSearchArticle)
		// 获取网站信息
		webGroup.POST("/website/info",web_controller.PostWebsiteInfo)
		// 获取所有友情链接
		webGroup.POST("/links",web_controller.PostLinks)
		// 需要登陆权限
		webAuthGroup := webGroup.Group("/auth")
		webAuthGroup.Use(middleware.WebAuth())
		{
			webAuthGroup.POST("/personal/upload/avatar",web_controller.AuthPostPersonalUploadAvatar)
			webAuthGroup.POST("/personal/edit",web_controller.AuthPostPersonalEdit)
		}
	}
	// 后台接口
	// 自由访问
	backendGroup := app.Group("/backend")
	{
		// 图形验证码
		backendGroup.GET("/captcha/img",backend_controller.GetImgCaptcha)
		// 注册
		backendGroup.POST("/manager/register",backend_controller.PostRegisterManager)
		// 登录
		backendGroup.POST("/manager/login",backend_controller.PostLoginManager)
		// 获取网站信息
		backendGroup.POST("/website/info",backend_controller.PostWebsiteInfo)
		// 服务器信息
		backendGroup.POST("/server/info",backend_controller.PostServerInfo)
		// 需要权限
		backAuthGroup := backendGroup.Group("/auth")
		backAuthGroup.Use(middleware.BackendAuth())
		{
			// 获取后台菜单
			backAuthGroup.POST("/menu/list",backend_controller.PostBackendMenu)
			// 获取前台菜单
			backAuthGroup.POST("/menu/web/list",backend_controller.PostWebMenu)
			// 修改前台菜单
			backAuthGroup.POST("/menu/web/edit",backend_controller.PostEditWebMenu)
			// 添加前台菜单
			backAuthGroup.POST("/menu/web/add",backend_controller.PostAddWebMenu)
			// 删除前台菜单
			backAuthGroup.POST("/menu/web/remove",backend_controller.PostRemoveMenu)
			// 获取所有文章
			backAuthGroup.POST("/article/all",backend_controller.PostArticleAll)
			// 批量删除文章
			backAuthGroup.POST("/article/remove/arr",backend_controller.PostRemoveArticleArr)
			// 上传文章banner
			backAuthGroup.POST("/article/upload/image/banner",backend_controller.PostArticleUploadBanner)
			// 上传文章封面图片
			backAuthGroup.POST("/article/upload/image/cover",backend_controller.PostArticleUploadCover)
			// 上传文章视频
			backAuthGroup.POST("/article/upload/video",backend_controller.PostArticleUploadVideo)
			// 添加文章
			backAuthGroup.POST("/article/add",backend_controller.PostArticleAdd)
			// 根据ID 获取文章数据
			backAuthGroup.POST("/article/info",backend_controller.PostArticleInfo)
			// 更具ID 更新文章
			backAuthGroup.POST("/article/edit",backend_controller.PostArticleEdit)
			// 富文本编辑器上传图片
			backAuthGroup.POST("/editor/upload/img",backend_controller.PostEditorUploadImg)
			// 获取用户列表
			backAuthGroup.POST("/consumer/list",backend_controller.PostConsumerList)
			// 修改用户状态
			backAuthGroup.POST("/consumer/edit/state",backend_controller.PostConsumerState)
			// 获取所有评论数据
			backAuthGroup.POST("/comment/list",backend_controller.PostCommentAll)
			// 删除评论
			backAuthGroup.POST("/comment/remove",backend_controller.PostRemoveComment)
			// 获取所有链接
			backAuthGroup.POST("/link/list",backend_controller.PostLinkAll)
			// 修改链接文本
			backAuthGroup.POST("/link/edit/text",backend_controller.PostLinkEdit)
			// 更新链接图片
			backAuthGroup.POST("/link/icon/edit",backend_controller.PostLinkIconEdit)
			// 删除链接
			backAuthGroup.POST("/link/remove",backend_controller.PostLinkRemoveById)
			// 添加链接
			backAuthGroup.POST("/link/insert/one",backend_controller.PostLinkInsertOne)
			// 更新网站头像图片
			backAuthGroup.POST("/website/avatar/edit",backend_controller.PostWebsiteAvatarEdit)
			// 更新网站LOGO
			backAuthGroup.POST("/website/logo/edit",backend_controller.PostWebsiteLogoEdit)
			// 更新网站Banner
			backAuthGroup.POST("/website/banner/edit",backend_controller.PostWebsiteBannerEdit)
			// 更新网站微博二维码
			backAuthGroup.POST("/website/webo/edit",backend_controller.PostWebsiteWeboEdit)
			// 更新网站微信二维码
			backAuthGroup.POST("/website/wechar/edit",backend_controller.PostWebsiteWecharEdit)
			// 更新网站基本信息
			backAuthGroup.POST("/website/info/edit",backend_controller.PostEditWebsiteInfo)
			// 更新关于我们
			backAuthGroup.POST("/website/about/edit",backend_controller.PostEditWebsiteAbout)
			// 更新用户协议
			backAuthGroup.POST("/website/agreement/edit",backend_controller.PostEditWebsiteAgreement)
			// 更新隐私声明
			backAuthGroup.POST("/website/privacy/edit",backend_controller.PostEditWebsitePrivacy)
			// 更新APK
			backAuthGroup.POST("/website/apk/edit",backend_controller.PostWebsiteApkEdit)
		}
	}
	return app
}
