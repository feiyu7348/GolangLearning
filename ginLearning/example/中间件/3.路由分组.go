// @author:zfy
// @date:2023/8/12 10:49

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserListView(c *gin.Context) {
	var userList = []UserInfo{
		{"zfy", 27},
		{"whh", 25},
	}

	c.JSON(http.StatusOK, Response{0, userList, "请求成功"})
}

func ArticleListView(c *gin.Context) {
	var articleList = []article{
		{"go", "123"},
		{"python", "123"},
	}

	c.JSON(http.StatusOK, Response{1, articleList, "请求成功"})
}

func UserRouter(router *gin.RouterGroup) {
	userManager := router.Group("user_manager")
	{
		userManager.GET("/users", UserListView)
	}
}

func ArticleRouter(router *gin.RouterGroup) {
	articleManager := router.Group("user_article")
	{
		articleManager.GET("/articles", ArticleListView)
	}
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	UserRouter(api)
	ArticleRouter(api)

	err := router.Run(":8083")
	if err != nil {
		return
	}
}
