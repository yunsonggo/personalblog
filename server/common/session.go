package common

import (
	"2021/yunsongcailu/yunsong_server/config"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//初始化
func InitSession(engine *gin.Engine) {
	var sessConf = config.Conf.Session
	var resConf = config.Conf.Redis
	store, err := redis.NewStore(10, "tcp", resConf.RedisAddr, resConf.RedisPwd, []byte(sessConf.SessionSecret))
	if err != nil {
		panic(fmt.Sprintf("初始化sessionStore错误:%v\n", err))
	}
	engine.Use(sessions.Sessions(sessConf.SessionName, store))
}

//设置  session
func SetSess(ctx *gin.Context, key string, value string,theMaxAge int) error {
	session := sessions.Default(ctx)
	if session == nil {
		err := errors.New("实例化session错误")
		panic(fmt.Sprintf("设置session失败:%v\n", err))

	}
	if theMaxAge > 0  {
		option := sessions.Options{
			MaxAge: theMaxAge,
		}
		session.Options(option)
	}
	session.Set(key, value)
	return session.Save()
}

//获取 session
func GetSess(ctx *gin.Context, key string) interface{} {
	session := sessions.Default(ctx)
	if session == nil {
		err := errors.New("实例化session错误")
		panic(fmt.Sprintf("获取session失败:%v\n", err))

	}
	return session.Get(key)
}
