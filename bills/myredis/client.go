package myredis

import (
	"github.com/go-redis/redis/v7"
	"log"
)

type MyRedis struct {
	Client *redis.Client
}

func NewMyRedis(dsn , pass string) MyRedis {
	client := redis.NewClient(&redis.Options{
		Addr: dsn, Password: pass, 
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Panicf("%v : %v  --- %v ", err, dsn, dsn )
	}
	return MyRedis{
		Client: client,
	}
}
