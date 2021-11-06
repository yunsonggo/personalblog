package dial

import (
	"github.com/go-redis/redis/v7"
)

var RedisDB *redis.Client

func InitRedisClient(addr string,password string) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB:1,
	})
	RedisDB = client
	return
}