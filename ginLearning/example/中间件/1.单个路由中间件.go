// @author:zfy
// @date:2023/8/12 10:06

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1")
	c.JSON(http.StatusOK, gin.H{"msg": "m1"})
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
		c.Abort() // 后面的就不相应了
	}, func(c *gin.Context) {
		fmt.Println(1)
	}, func(c *gin.Context) {
		fmt.Println(2)
	}, func(c *gin.Context) {
		fmt.Println(3)
	}, m1)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
