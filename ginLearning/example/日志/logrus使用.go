// @author:zfy
// @date:2023/8/12 14:30

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Error("出错了")
	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {})

	err := router.Run(":8083")
	if err != nil {
		return
	}
}
