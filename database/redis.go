package database

import (
	"context"
	"log"
	"sync"
)

var defaultClient *redis.Client
var once sync.Once

func GetDefaultClient() *redis.Client {
	once.Do(func() {
		defaultClient = redis.NewClient(&redis.Options{
			//Addr:     gconf.GConf.RedisHostPort,
			//Password: gconf.GConf.RedisPwd,
		})
		pong, err := defaultClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("redis connect ping failed, err:%s", err.Error())
		} else {
			log.Println(pong)
		}
	})
	return defaultClient
}
