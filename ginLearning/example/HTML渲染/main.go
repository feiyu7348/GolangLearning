// author:zfy  date:2022/9/25

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template.html", gin.H{"title": "Main html"})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
