package config

import "2021/yunsongcailu/yunsong_server/dial"

func InitRrsEngine() {
	mc := Conf.Mysql
	rc := Conf.Redis
	dial.MysqlEngine(mc.MysqlAddr,mc.MysqlDriver,mc.MysqlShow)
	dial.InitRedisClient(rc.RedisAddr,rc.RedisPwd)
}
