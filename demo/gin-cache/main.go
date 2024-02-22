package main

import (
	"fmt"
	cache "github.com/chenyahui/gin-cache"
	"log"
	"time"

	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	app.GET("/hello", cache.CacheByRequestURI(
		memoryStore,
		2*time.Second,
		//cache.WithOnHitCache(mid2),
		cache.WithOnShareSingleFlight(mid3),
	), handler)
	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}

func handler(c *gin.Context) {
	key := c.DefaultQuery("key", "default")
	log.Println("func handler", key)
	c.String(200, "hello world")
}

func mid2(c *gin.Context) {

	fmt.Println("func mid2")
}

func mid3(c *gin.Context) {
	fmt.Println("func mid3")

}
