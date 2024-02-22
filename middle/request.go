package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"log"
	"sort"
	"sync"
)

func RateLimit() func(c *gin.Context) {
	//每秒往桶里放100,桶大小为100
	limiter := rate.NewLimiter(100, 100)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			fmt.Println("请求被限流:", c.Request.URL.Path)
			c.Abort()
			c.JSON(200, gin.H{"msg": "系统繁忙", "data": []interface{}{}, "total": 0, "code": 400})
			return
		}
		c.Next()
	}
}

//只启动一个pod TODO panic之后 后续请求都会 报请勿重复提交
var requestMap sync.Map

func ReSubmitCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		adminId := c.GetInt("admin")
		var key = fmt.Sprintf("%d%s", adminId, c.Request.URL.Path)
		_, loaded := requestMap.LoadOrStore(key, 1)
		if loaded {
			c.Abort()
			c.JSON(200, gin.H{"msg": "请勿重复提交", "code": 400})
			return
		}
		c.Next()
		requestMap.LoadAndDelete(key)
	}
}

//请求白名单地址
func WhiteList(ips ...string) func(c *gin.Context) {
	sort.Strings(ips)
	return func(c *gin.Context) {
		ip := c.ClientIP()
		_, isExist := util.IsExistString(ips, ip)
		if !isExist {
			log.Println("非法请求ip:", ip)
			c.Abort()
			c.JSON(200, gin.H{"msg": "非法IP", "code": 400})
			return
		}
		c.Next()
	}
}
