package main

import (
	"gin-trans/trans"
	"net/http"

	"github.com/gin-gonic/gin"
)

type YourStruct struct {
	YourField string `json:"your_field" binding:"required,alphanum" label:"你的字段"`
}

func main() {
	r := gin.Default()
	trans.InitTrans()
	r.POST("/validate", func(c *gin.Context) {
		var requestData YourStruct
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": trans.Error(err)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Input is valid"})
	})
	r.Run(":8080")
}
