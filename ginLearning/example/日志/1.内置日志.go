// @author:zfy
// @date:2023/8/12 14:11

package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("./example/日志/gin.log")
	gin.DefaultWriter = io.MultiWriter(file)

	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {})
	router.POST("/post", func(c *gin.Context) {})

	err := router.Run(":8083")
	if err != nil {
		return
	}
}
