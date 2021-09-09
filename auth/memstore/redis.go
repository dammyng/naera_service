package memstore

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type Redis struct {
	Client *redis.Client
}

func NewRedisClient(dsn, pass string) Redis {
	client := redis.NewClient(&redis.Options{
		Addr: dsn, Password: pass,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return Redis{
		Client: client,
	}
}

func (r Redis) Get(arg interface{}) (string, error) {
	key := fmt.Sprintf("%v", arg)
	return r.Client.Get(key).Result()
}

func (r Redis) Set(key, value string, duration time.Duration) error {
	return nil
}
