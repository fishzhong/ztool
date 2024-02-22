package main

import (
	"fmt"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"time"
	"webserver/handler"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	r.GET("/game", cache.CacheByRequestURI(
		memoryStore,
		2*time.Second,
		//cache.WithOnHitCache(mid2),
		//cache.WithOnShareSingleFlight(mid3),
	), handler.Handler)
	r.GET("/game2", handler.Handler2)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func mid2(c *gin.Context) {
	fmt.Println("func mid2")
}

func mid3(c *gin.Context) {
	fmt.Println("func mid3")

}
