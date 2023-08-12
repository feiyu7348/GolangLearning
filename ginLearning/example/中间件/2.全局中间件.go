// @author:zfy
// @date:2023/8/12 10:26

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func m2(c *gin.Context) {
	fmt.Println("m2 ...in")
	c.Set("name", User{"枫枫", 21})
	c.Set("user", User{Name: "zfy", Age: 27})
	fmt.Println("m2 ...out")
}

func main() {
	router := gin.Default()

	router.Use(m2)

	router.GET("/", func(c *gin.Context) {
		name, _ := c.Get("name")
		fmt.Println(name)
		_user, _ := c.Get("user")
		user := _user.(User)
		fmt.Println(user.Name, user.Age)

		c.JSON(http.StatusOK, gin.H{"msg": user})
	})

	err := router.Run(":8083")
	if err != nil {
		return
	}
}
