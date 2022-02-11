package redisconnet

import (
	"github.com/go-redis/redis"
)

var rdc *redis.Client

func InitClient() (err error) {
	rdc = redis.NewClient(&redis.Options{
		Addr:     "47.93.125.10:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdc.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
