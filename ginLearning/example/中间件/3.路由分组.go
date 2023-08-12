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

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserListView(c *gin.Context) {
	var userList []UserInfo = []UserInfo{
		{"zfy", 27},
		{"whh", 25},
	}

	c.JSON(http.StatusOK, Response{0, userList, "请求成功"})
}

func main() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/users", UserListView)

	err := router.Run(":8083")
	if err != nil {
		return
	}
}
