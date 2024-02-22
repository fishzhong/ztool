package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"path"
	"sync/atomic"
	"time"
)

var count int64

func main() {
	// 1.创建路由
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	//go func() {
	//	time.Sleep(10 * time.Second)
	//	os.Exit(1)
	//}()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	route := r.Group("admin").Use(middle1())
	route.GET("test", handler)
	route.Use(middle2(), middle4())
	route.GET("test2", middle3(), middle5(), handler)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

func UserFileDownloadCommonService(c *gin.Context) {
	c.RemoteIP()
	filePath := c.Query("url")
	//打开文件
	fileTmp, errByOpenFile := os.Open(filePath)
	if errByOpenFile != nil {
		//log.Error("获取文件失败")
		c.String(http.StatusOK, errByOpenFile.Error())
		return
	}
	defer fileTmp.Close()
	//获取文件的名称
	fileName := path.Base(filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	if !PathExists(filePath) || !PathExists(fileName) {
		c.String(http.StatusOK, "没有查询到")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
	return
}
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func handler(c *gin.Context) {

	fmt.Println("func handler ")
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

func middle() func(c *gin.Context) {
	limit := rate.Every(time.Millisecond)
	limiter := rate.NewLimiter(limit, 100)
	return func(c *gin.Context) {
		if limiter.AllowN(time.Now(), 1) {
			fmt.Println("event allowed")
		} else {
			atomic.AddInt64(&count, 1)
			fmt.Println("event not allowed,合计", count)
		}
		c.Next()
	}
}

func middle1() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("func middle1")
		c.Next()
	}
}

func middle2() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("func middle2")
		c.Next()
	}
}
func middle3() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("func middle3")
		c.Next()
	}
}

func middle4() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("func middle4")
		c.Next()
	}
}

func middle5() func(c *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("func middle5")
		c.Next()
	}
}
