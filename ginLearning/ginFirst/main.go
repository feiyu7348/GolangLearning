// author:zfy  date:2022/9/24

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong/main",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	err := r.Run(":8083")
	if err != nil {
		return 
	}
}


