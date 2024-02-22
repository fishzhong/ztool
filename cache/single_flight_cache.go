package cache

import (
	"context"
	"errors"
	"golang.org/x/sync/singleflight"
	"time"
)

var ErrCacheMiss = errors.New("cache miss")
var g singleflight.Group

type LoadDataFunc func() (interface{}, error)

func GetData(ctx context.Context, key string, LoadFromCache, LoadFormDb LoadDataFunc) (interface{}, error) {
	data, err := LoadFromCache()
	if err != nil && errors.Is(err, ErrCacheMiss) {
		// 使用 DoChan 结合 select 做超时控制
		result := g.DoChan(key, func() (interface{}, error) {
			go func() {
				time.Sleep(100 * time.Millisecond)
				//fmt.Printf("Deleting key: %v\n", key)
				g.Forget(key)
			}()
			data, err = LoadFormDb()
			return data, nil
		})
		select {
		case r := <-result:
			//fmt.Println(r.Shared)
			return r.Val, r.Err
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	return data, nil
}
