// @author:zfy
// @date:2023/8/12 14:01

package main

import (
	"github.com/gin-gonic/gin"
)

func JwtTokenMiddleware(c *gin.Context) {
	// 获取请求头的token
	token := c.GetHeader("token")
	// 调用jwt的验证函数
	if token == "1234" {
		// 验证通过
		c.Next()
		return
	}
	// 验证不通过
	c.JSON(200, gin.H{"msg": "权限验证失败"})
	c.Abort()
}

func main() {
	router := gin.Default()

	api := router.Group("/api")

	apiUser := api.Group("")
	{
		apiUser.POST("login", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "登录成功"})
		})
	}
	apiHome := api.Group("system").Use(JwtTokenMiddleware)
	{
		apiHome.GET("/index", func(c *gin.Context) {
			c.String(200, "index")
		})
		apiHome.GET("/home", func(c *gin.Context) {
			c.String(200, "home")
		})
	}

	err := router.Run(":8083")
	if err != nil {
		return
	}
}
