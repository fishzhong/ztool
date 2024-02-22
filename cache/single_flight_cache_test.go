package cache

import (
	"context"
	"testing"
)

func BenchmarkGetData(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		GetData(ctx, "strconv.Itoa(i)", loadFromCache, loadFromDb)
	}
}

func loadFromCache() (interface{}, error) {

	return nil, ErrCacheMiss
}
func loadFromDb() (interface{}, error) {
	//time.Sleep(time.Second)
	return nil, nil
}
