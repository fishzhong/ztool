package database

import (
	goredislib "github.com/go-redis/redis/v9"
	"log"
	"sync"
	"sync/atomic"
)

var rs *redsync.Redsync
var mu sync.Mutex
var done uint32

func GetDLocks() *redsync.Redsync {
	if atomic.LoadUint32(&done) == 0 {
		mu.Lock()
		defer mu.Unlock()
		if done == 0 {
			defer atomic.StoreUint32(&done, 1)
			client := goredislib.NewClient(&goredislib.Options{
				//Addr:     gconf.GConf.RedisHostPort,
				//Password: gconf.GConf.RedisPwd,
			})
			pool := goredis.NewPool(client)
			rs = redsync.New(pool)
			log.Println("初始化")
		}
	}
	return rs
}
